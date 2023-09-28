package errors

import (
	"fmt"
	"movie_booking/models"
)

type SeatsNotFoundError struct {
	Id models.AudiIdModel
}

func (err *SeatsNotFoundError) Error() string {
	return fmt.Sprintf("seats for audi id: %d not found.", err.Id)
}
