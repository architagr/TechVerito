package service

import "movie_booking/models"

type swatchhBharatTaxService struct {
	name string
	rate float32
}

func (taxSvc *swatchhBharatTaxService) Calculate(amount *models.AmountModel) {
	if len(amount.Taxes) == 0 {
		amount.Taxes = make([]models.TaxModel, 0)
	}
	amount.Taxes = append(amount.Taxes, models.TaxModel{
		Name:            taxSvc.name,
		Rate:            taxSvc.rate,
		ChargableAmount: amount.SubTotal * taxSvc.rate / 100,
	})
}

func InitSwatchhBharatTaxService(rate float32) ITaxService {
	return &swatchhBharatTaxService{
		name: "Swatch Bharat Tax",
		rate: rate,
	}
}
