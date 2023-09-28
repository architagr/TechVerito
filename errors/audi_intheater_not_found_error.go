package errors

import (
	"fmt"
	"movie_booking/models"
)

type AudiInTheaterNotFoundError struct {
	Id models.TheaterIdModel
}

func (err *AudiInTheaterNotFoundError) Error() string {
	return fmt.Sprintf("audi in theater with id: %d not found.", err.Id)
}
