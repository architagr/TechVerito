package presentation

import (
	"movie_booking/enums"
	"movie_booking/service"
)

var salesService service.ISalesService
var showSeatService service.IShowSeatService
var showService service.IShowService
var ticketService service.ITicketService

func initAllServices() {
	allTaxes := make([]service.ITaxService, 0, 3)
	allTaxes = append(allTaxes, service.TaxFactory(enums.TAX_KRISHI_KALYAN), service.TaxFactory(enums.TAX_SERVICE), service.TaxFactory(enums.TAX_SWATCHH_BHARAT))
	showSeatService = service.InitShowSeatService(showRepo, showSeatRepo, seatTypeRepo, showSeatTypeRepo, seatRepo)
	showService = service.InitShowService(showRepo, theaterRepo, audiRepo, movieRepo, showSeatService)
	salesService = service.InitSalesService(ticketRepo)
	ticketService = service.InitTicketService(ticketRepo, showSeatService, allTaxes)
}
