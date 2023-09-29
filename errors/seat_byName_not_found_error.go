package errors

import (
	"fmt"
	"movie_booking/models"
)

type SeatsByNameNotFoundError struct {
	AudiId models.AudiIdModel
	Name   string
}

func (err *SeatsByNameNotFoundError) Error() string {
	return fmt.Sprintf("seat(%s) for audi id: %d not found.", err.Name, err.AudiId)
}
