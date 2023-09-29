package service

import (
	"movie_booking/enums"
	"movie_booking/models"
)

type ITaxService interface {
	Calculate(amount *models.AmountModel)
}

func TaxFactory(taxType enums.TaxEnum) ITaxService {
	if taxType == enums.TAX_KRISHI_KALYAN {
		return InitKrishiKalyanTaxService(.5)
	} else if taxType == enums.TAX_SWATCHH_BHARAT {
		return InitSwatchhBharatTaxService(.5)
	} else if taxType == enums.TAX_SERVICE {
		return InitServiceTaxService(14)
	}
	return nil
}
