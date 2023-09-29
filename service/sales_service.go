package service

import (
	"movie_booking/models"
	"movie_booking/persistance"
)

type ISalesService interface {
	GetTotalSalesReport() models.TotalSalesResponse
}

func InitSalesService(ticketRepo persistance.ITicketPersistance) ISalesService {
	return &SalesService{
		ticketRepo: ticketRepo,
	}
}

type SalesService struct {
	ticketRepo persistance.ITicketPersistance
}

func (salesSvc *SalesService) GetTotalSalesReport() models.TotalSalesResponse {
	colectedTaxes := make(map[string]float32)
	revenue := float32(0)

	allTickets := salesSvc.ticketRepo.GetAll()
	for _, ticket := range allTickets {
		revenue += ticket.Charges.SubTotal
		for _, tax := range ticket.Charges.Taxes {
			taxAmount := colectedTaxes[tax.Name]
			taxAmount += tax.ChargableAmount
			colectedTaxes[tax.Name] = taxAmount
		}
	}

	return models.TotalSalesResponse{
		Revenue:       revenue,
		ColectedTaxes: colectedTaxes,
	}
}
