package persistance

import "movie_booking/models"

type ITicketPersistance interface {
	GetAll() []models.TicketModel
	Add(ticket models.TicketModel) models.TicketModel
}
type ticketPersistance struct {
	lastId  int
	tickets []models.TicketModel
}

func (repo *ticketPersistance) GetAll() []models.TicketModel {
	return repo.tickets
}
func (repo *ticketPersistance) Add(ticket models.TicketModel) models.TicketModel {
	repo.lastId++
	ticket.Id = models.TicketIdModels(repo.lastId)
	repo.tickets = append(repo.tickets, ticket)
	return ticket
}

func InitTicketPersistance() ITicketPersistance {
	return &ticketPersistance{
		lastId:  0,
		tickets: make([]models.TicketModel, 0),
	}
}
