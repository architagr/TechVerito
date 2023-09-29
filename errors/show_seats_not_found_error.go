package errors

import (
	"fmt"
	"movie_booking/models"
)

type ShowSeatsNotFoundError struct {
	ShowId models.ShowIdModel
}

func (err *ShowSeatsNotFoundError) Error() string {
	return fmt.Sprintf("seats for show id: %d not found.", err.ShowId)
}
