package models

type AmountModel struct {
	SubTotal float32    `json:"subTotal"`
	Taxes    []TaxModel `json:"taxes"`
}

func (bill *AmountModel) GetTotal() (total float32) {
	total = bill.SubTotal
	for _, tax := range bill.Taxes {
		total += tax.ChargableAmount
	}
	return
}
