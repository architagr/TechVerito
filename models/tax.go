package models

import "fmt"

type TaxModel struct {
	Name            string  `json:"name"`
	Rate            float32 `json:"rate"`
	ChargableAmount float32 `json:"chargableAmount"`
}

func (tax *TaxModel) ToString() string {
	return fmt.Sprintf("%s @%.2f %%: Rs. %.2f\n", tax.Name, tax.Rate, tax.ChargableAmount)
}
