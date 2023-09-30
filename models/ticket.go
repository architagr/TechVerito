package models

import (
	"fmt"
	"movie_booking/enums"
	"time"
)

type TicketIdModels int
type TicketModel struct {
	Id            TicketIdModels         `json:"id"`
	Charges       AmountModel            `json:"charges"`
	ShowId        ShowIdModel            `json:"showId"`
	SeatIds       []SeatIdModel          `json:"seatIds"`
	TimeOfBooking time.Time              `json:"timeOfBooking"`
	Status        enums.TicketStatusEnum `json:"status"`
}

func (tkt *TicketModel) ToString() string {
	return fmt.Sprintf("Successfully Booked - Show %d\n%s", tkt.ShowId, tkt.Charges.ToString())
}
