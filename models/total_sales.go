package models

type TotalSalesResponse struct {
	Revenue       float32            `json:"revenue"`
	ColectedTaxes map[string]float32 `json:"collectedTaxes"`
}
