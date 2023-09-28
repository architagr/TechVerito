package errors

import (
	"fmt"
	"movie_booking/models"
)

type TheaterNotFoundError struct {
	Id models.TheaterIdModel
}

func (err *TheaterNotFoundError) Error() string {
	return fmt.Sprintf("theater by id: %d not found.", err.Id)
}
