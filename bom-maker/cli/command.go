package cli

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
	"unicode/utf8"

	"github.com/bom-maker/bomcsv"
	"github.com/bom-maker/mouser"
	"github.com/bom-maker/mouser/model"
	cli "github.com/jawher/mow.cli"
)

var (
	// MouserAPIKey is the key used to make API calls
	MouserAPIKey *string
)

// Process processes command line args
func Process(appName, appDesc, appVersion string) {
	app := cli.App(appName, appDesc)
	app.Spec = "-k"

	app.Version("v version", fmt.Sprintf("%s version %s", appName, appVersion))

	MouserAPIKey = app.StringOpt("k mouser-api-key", "", "Mouser API key to use to get parts' information")

	app.Command("generate", "Generate a BOM file", generate)

	app.Action = func() {
		fmt.Println("Please choose a command arg!")
	}

	app.Run(os.Args)
}

func generate(cmd *cli.Cmd) {
	cmd.Spec = "-i -o [-s] [-t]"

	input := cmd.StringOpt("i input", "", "Input file")
	output := cmd.StringOpt("o output", "", "Output file")
	separator := cmd.StringOpt("s sep", ";", "Separator used to read the input file")
	outputSeparator := cmd.StringOpt("t outputSep", ";", "Separator used to write the output file")

	cmd.Action = func() {
		fmt.Println(input, output, separator, outputSeparator)

		// Reading input file
		file, err := os.Open(*input)
		if err != nil {
			fmt.Printf("Unable to read file, err=%v", err)
			os.Exit(1)
		}
		defer file.Close()
		reader := csv.NewReader(file)
		reader.Comma, _ = utf8.DecodeRune([]byte(*separator))

		// Reading header
		headers, err := reader.Read()
		if err != nil {
			fmt.Printf("Unable to read record from CSV, err=%+v", err)
		}

		csvParts := make([]bomcsv.Part, 0)
		api := mouser.NewAPI(*MouserAPIKey)
		for {
			record, err := reader.Read()
			if err != nil && err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Unable to read record from CSV, err=%+v\n", err)
			}

			// Reading all fields and create a part from it
			csvPart := bomcsv.Part{}
			for k, v := range record {
				err := csvPart.SetPartField(headers[k], v)
				if err != nil {
					// Silently ignoring errors
					continue
				}
			}

			if csvPart.MouserRef != "" {
				csvParts = append(csvParts, csvPart)
			} else {
				fmt.Printf("Warning: part %s (%s) has an empty MouserRef, skipping.\n", csvPart.Device, csvPart.Parts)
			}
		}

		var wg sync.WaitGroup
		parts := make([]model.Part, 0)
		mutex := sync.Mutex{}
		for _, v := range csvParts {
			fmt.Printf("Calling Mouser API for part(s) %s (Device=%s, MouserRef=%s)\n", v.Parts, v.Device, v.MouserRef)
			wg.Add(1)
			go func(wg *sync.WaitGroup, mouserRef string) {
				defer wg.Done()
				p, err := api.SearchByPartNumber(mouserRef)
				if err != nil {
					fmt.Errorf("An error occurred looking for MouserRef %s, err=%+v", mouserRef, err)
				} else {
					mutex.Lock()
					parts = append(parts, *p)
					mutex.Unlock()
				}
			}(&wg, v.MouserRef)
		} // for

		wg.Wait() // Wait for all go routines to finish

		fmt.Println("==========")
		fmt.Printf("All parts:\n%+v\n", parts)
	}
}
