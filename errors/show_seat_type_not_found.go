package errors

import (
	"fmt"
	"movie_booking/models"
)

type ShowSeatTypeNotFoundError struct {
	ShowId models.ShowIdModel
}

func (err *ShowSeatTypeNotFoundError) Error() string {
	return fmt.Sprintf("show seat type data for show id : %d, not found", err.ShowId)
}
