package dto

type CarRequest struct {
	Name         string `json:"name"`
	Year         string `json:"year"`
	Manufacturer string `json:"manufacturer"`
	Description  string `json:"description"`
	Code         string `json:"code"`
	Collection   string `json:"collection"`
	BarCode      string `json:"barCode"`
	Category     string `json:"category"`
}
