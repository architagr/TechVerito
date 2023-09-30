package models

import "fmt"

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

func (bill *AmountModel) ToString() string {
	taxStr := ""

	for _, tax := range bill.Taxes {
		taxStr += tax.ToString()
	}

	return fmt.Sprintf("Subtotal: Rs. %.2f\n%s\nTotal: Rs.%.2f", bill.SubTotal, taxStr, bill.GetTotal())
}
