package bomcsv

import (
	"fmt"
	"strconv"
)

// Part represents an electronic part parsed from a CSV format
type Part struct {
	// Quantity is the number of this part in the BOM
	Quantity uint
	// Value is the value of the part (Farad, Ohms, etc.)
	Value string
	// Device is the device name used in the CAD library
	Device string
	// MouserRef is the reference used from Mouser catalog
	MouserRef string
	// Package is the package name of the part
	Package string
	// Parts' is the name of the part used in the schematic
	Parts string
	// Description is the associated description of the part
	Description string
}

// SetPartField set the part's field corresponding to the given header's name and associated value
func (p *Part) SetPartField(header, value string) error {
	switch header {
	case "Qty":
		i, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		p.Quantity = uint(i)
	case "Value":
		p.Value = value
	case "Device":
		p.Device = value
	case "MouserRef":
		p.MouserRef = value
	case "Package":
		p.Package = value
	case "Parts":
		p.Parts = value
	case "Description":
		p.Description = value
	default:
		return fmt.Errorf("Ignoring unknown header %s", header)
	}

	return nil
}
