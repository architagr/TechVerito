package errors

import (
	"fmt"
	"movie_booking/models"
)

type ShowNotFoundError struct {
	Id models.ShowIdModel
}

func (err *ShowNotFoundError) Error() string {
	return fmt.Sprintf("shows with id: %d not found", err.Id)
}
