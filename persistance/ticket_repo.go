package persistance

import "movie_booking/models"

type ITicketPersistance interface {
	GetAll() []models.TicketModel
	Add(ticket models.TicketModel)
}
type TicketPersistance struct {
	lastId  int
	tickets []models.TicketModel
}

func (repo *TicketPersistance) GetAll() []models.TicketModel {
	return repo.tickets
}
func (repo *TicketPersistance) Add(ticket models.TicketModel) {
	repo.lastId++
	ticket.Id = models.TicketIdModels(repo.lastId)
	repo.tickets = append(repo.tickets, ticket)
}

func InitTicketPersistance() ITicketPersistance {
	return &TicketPersistance{
		lastId:  0,
		tickets: make([]models.TicketModel, 0),
	}
}
