package errors

import (
	"fmt"
	"movie_booking/models"
)

type CityNotFoundError struct {
	Id models.CityIdModel
}

func (err *CityNotFoundError) Error() string {
	return fmt.Sprintf("city by id: %d not found.", err.Id)
}
