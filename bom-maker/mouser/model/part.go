package model

import (
	"encoding/json"
	"strconv"
)

// APIUInt is an unsigned integer number represented in the Mouser API
type APIUInt uint

// APIFloat32 is a float32 represented in the Mouser API
type APIFloat32 float32

// Part represents an electronic component from Mouser perspective
type Part struct {
	Availability           string  `json:"Availability"`
	DatasheetURL           string  `json:"DataSheetUrl"`
	Description            string  `json:"Description"`
	FactoryStock           APIUInt `json:"FactoryStock"`
	ImagePath              string  `json:"ImagePath"`
	Category               string  `json:"Category"`
	LeadTime               string  `json:"LeadTime"`
	LifecycleStatus        string  `json:"LifecycleStatus"`
	Manufacturer           string  `json:"Manufacturer"`
	ManufacturerPartNumber string  `json:"ManufacturerPartNumber"`
	Min                    APIUInt `json:"Min"`
	Mult                   APIUInt `json:"Mult"`
	MouserPartNumber       string  `json:"MouserPartNumber"`
	ProductAttributes      []struct {
		AttributeName  string `json:"AttributeName"`
		AttributeValue string `json:"AttributeValue"`
	} `json:"ProductAttributes"`
	PriceBreaks []struct {
		Quantity APIUInt `json:"Quantity"`
		Price    string  `json:"Price"`
		Currency string  `json:"Currency"`
	} `json:"PriceBreaks,omitempty"`
	AlternatePackagings []struct {
		APMfrPN string `json:"APMfrPN"`
	} `json:"AlternatePackagings"`
	ProductDetailURL     string  `json:"ProductDetailUrl"`
	Reeling              bool    `json:"Reeling"`
	ROHSStatus           string  `json:"ROHSStatus"`
	SuggestedReplacement string  `json:"SuggestedReplacement"`
	MultiSimBlue         APIUInt `json:"MultiSimBlue"`
	UnitWeightKg         struct {
		UnitWeight APIFloat32 `json:"UnitWeight"`
	} `json:"UnitWeightKg"`
	StandardCost struct {
		StandardCost APIFloat32 `json:"Standardcost"`
	} `json:"StandardCost"`
	IsDiscontinued        string `json:"IsDiscontinued"`
	RTM                   string `json:"RTM"`
	MouserProductCategory string `json:"MouserProductCategory"`
	IPCCode               string `json:"IPCCode"`
	SField                string `json:"SField"`
	VNum                  string `json:"VNum"`
	ActualMfrName         string `json:"ActualMfrName"`
	AvailableOnOrder      string `json:"AvailableOnOrder"`
	RestrictionMessage    string `json:"RestrictionMessage"`
	PID                   string `json:"PID"`

	ProductCompliance []struct {
		ComplianceName  string `json:"ComplianceName"`
		ComplianceValue string `json:"ComplianceValue"`
	} `json:"ProductCompliance"`
}

// UnmarshalJSON unmarshals a APIUInt object type
func (i *APIUInt) UnmarshalJSON(data []byte) error {
	// Basically treat input data as string and convert it to go type
	// Parameters can be:
	//   - quoted
	//   - unquoted
	//   - empty quoted
	//   - null
	var dataStr string
	s, _ := strconv.Unquote(string(data))
	s = strconv.Quote(s)
	//log.Printf("APIUInt, s=%s", s)
	if err := json.Unmarshal([]byte(s), &dataStr); err != nil {
		return err
	}

	r, err := strconv.ParseUint(dataStr, 10, 32)
	if err != nil {
		*i = APIUInt(0)
	} else {
		*i = APIUInt(r)
	}

	return nil
}

// UnmarshalJSON unmarshals a APIFloat32 object type
func (f *APIFloat32) UnmarshalJSON(data []byte) error {
	// Basically treat input data as string and convert it to go type
	// Parameters can be:
	//   - quoted
	//   - unquoted
	//   - empty quoted
	//   - null
	var dataStr string
	s, _ := strconv.Unquote(string(data))
	s = strconv.Quote(s)
	//log.Printf("APIFloat32, s=%s", s)
	if err := json.Unmarshal([]byte(s), &dataStr); err != nil {
		return err
	}

	r, err := strconv.ParseFloat(dataStr, 32)
	if err != nil {
		*f = APIFloat32(0)
	} else {
		*f = APIFloat32(r)
	}

	return nil
}
