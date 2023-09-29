package errors

import (
	"fmt"
	"movie_booking/models"
)

type ShowSeatNotFoundError struct {
	ShowId models.ShowIdModel
	SeatId models.SeatIdModel
}

func (err *ShowSeatNotFoundError) Error() string {
	return fmt.Sprintf("seat (%d) for show id: %d not found.", err.SeatId, err.ShowId)
}
