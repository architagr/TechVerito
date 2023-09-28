package errors

import (
	"fmt"
	"movie_booking/models"
)

type AudiNotFoundError struct {
	Id models.AudiIdModel
}

func (err *AudiNotFoundError) Error() string {
	return fmt.Sprintf("audi by id: %d not found.", err.Id)
}
