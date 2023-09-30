package service

import (
	"movie_booking/enums"
	"movie_booking/models"
	"movie_booking/persistance"
	"time"
)

type ITicketService interface {
	BookTicket(showId models.ShowIdModel, seatName []string) (*models.TicketModel, error)
}

type ticketService struct {
	ticketRepo      persistance.ITicketPersistance
	showSeatService IShowSeatService
	taxes           []ITaxService
}

func InitTicketService(ticketRepo persistance.ITicketPersistance,
	showSeatService IShowSeatService, taxes []ITaxService) ITicketService {
	return &ticketService{
		ticketRepo:      ticketRepo,
		showSeatService: showSeatService,
		taxes:           taxes,
	}
}

func (ticketSvc *ticketService) BookTicket(showId models.ShowIdModel, seatName []string) (*models.TicketModel, error) {
	seatsInfo, err := ticketSvc.showSeatService.CheckAvailability(showId, seatName)
	if err != nil {
		return nil, err
	}
	seatIds := make([]models.SeatIdModel, 0, len(seatName))
	subTotal := float32(0)
	for _, seat := range seatsInfo {
		seatIds = append(seatIds, seat.SeatId)
		ticketSvc.showSeatService.UpdateSeatStaus(showId, seat.SeatId, enums.SEAT_STATUS_BOOKED)
		subTotal += seat.Rate
	}
	amount := models.AmountModel{
		SubTotal: subTotal,
	}
	for _, taxSvc := range ticketSvc.taxes {
		taxSvc.Calculate(&amount)
	}
	ticket := ticketSvc.ticketRepo.Add(models.TicketModel{
		ShowId:        showId,
		TimeOfBooking: time.Now(),
		Status:        enums.TICKET_STATUS_BOOKED,
		SeatIds:       seatIds,
		Charges:       amount,
	})
	return &ticket, nil
}
