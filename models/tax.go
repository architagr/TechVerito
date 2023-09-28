package models

type TaxModel struct {
	Name            string  `json:"name"`
	Rate            float32 `json:"rate"`
	ChargableAmount float32 `json:"chargableAmount"`
}
