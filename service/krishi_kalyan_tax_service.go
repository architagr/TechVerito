package service

import "movie_booking/models"

type krishiKalyanTaxService struct {
	name string
	rate float32
}

func (taxSvc *krishiKalyanTaxService) Calculate(amount *models.AmountModel) {
	if len(amount.Taxes) == 0 {
		amount.Taxes = make([]models.TaxModel, 0)
	}
	amount.Taxes = append(amount.Taxes, models.TaxModel{
		Name:            taxSvc.name,
		Rate:            taxSvc.rate,
		ChargableAmount: amount.SubTotal * taxSvc.rate / 100,
	})
}

func InitKrishiKalyanTaxService(rate float32) ITaxService {
	return &krishiKalyanTaxService{
		name: "Krishi Kalyan Tax",
		rate: rate,
	}
}
