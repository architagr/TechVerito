package errors

import (
	"fmt"
	"movie_booking/models"
)

type SeatTypeNotFoundError struct {
	Id models.SeatTypeIdModel
}

func (err *SeatTypeNotFoundError) Error() string {
	return fmt.Sprintf("seat type by id: %d not found.", err.Id)
}
