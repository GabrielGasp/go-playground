package main

import (
	"encoding/json"
	"fmt"
)

// type AllowedParams struct {
// 	UtmSource   string `json:"utm_source,omitempty"`
// 	UtmMedium   string `json:"utm_medium,omitempty"`
// 	UtmCampaign string `json:"utm_campaign,omitempty"`
// 	Afp1        string `json:"afp1,omitempty"`
// 	Afp2        string `json:"afp2,omitempty"`
// 	Afp3        string `json:"afp3,omitempty"`
// 	Afp4        string `json:"afp4,omitempty"`
// 	Afp5        string `json:"afp5,omitempty"`
// 	Afp6        string `json:"afp6,omitempty"`
// 	Afp7        string `json:"afp7,omitempty"`
// 	Afp8        string `json:"afp8,omitempty"`
// 	Afp9        string `json:"afp9,omitempty"`
// 	Afp10       string `json:"afp10,omitempty"`
// }

var allowedParams = []string{
	"utm_source",
	"utm_medium",
	"utm_campaign",
	"afp1",
	"afp2",
	"afp3",
	"afp4",
	"afp5",
	"afp6",
	"afp7",
	"afp8",
	"afp9",
	"afp10",
}

func convertParamsToStruct(params map[string]string) map[string]string {
	// ap := AllowedParams{}

	// ap.UtmSource = params["utm_source"]
	// ap.UtmMedium = params["utm_medium"]
	// ap.UtmCampaign = params["utm_campaign"]
	// ap.Afp1 = params["afp1"]
	// ap.Afp2 = params["afp2"]
	// ap.Afp3 = params["afp3"]
	// ap.Afp4 = params["afp4"]
	// ap.Afp5 = params["afp5"]
	// ap.Afp6 = params["afp6"]
	// ap.Afp7 = params["afp7"]
	// ap.Afp8 = params["afp8"]
	// ap.Afp9 = params["afp9"]
	// ap.Afp10 = params["afp10"]

	// return ap

	ap := make(map[string]string)

	for _, v := range allowedParams {
		if val, ok := params[v]; ok {
			ap[v] = val
		}
	}

	return ap
}

func main() {
	m := map[string]string{
		"utm_source":   "google",
		"utm_medium":   "cpc",
		"utm_campaign": "campaign",
		"afp1":         "val1",
		"teste":        "teste",
	}

	ap := convertParamsToStruct(m)

	json, _ := json.Marshal(ap)

	fmt.Println(string(json))
}
