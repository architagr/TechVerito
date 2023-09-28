package errors

import (
	"fmt"
	"movie_booking/models"
)

type MovieNotFoundError struct {
	Id models.MovieIdModel
}

func (err *MovieNotFoundError) Error() string {
	return fmt.Sprintf("movie by id: %d not found.", err.Id)
}
