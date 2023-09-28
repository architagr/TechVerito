package errors

import (
	"fmt"
	"movie_booking/models"
)

type TheaterInCityNotFoundError struct {
	Id models.CityIdModel
}

func (err *TheaterInCityNotFoundError) Error() string {
	return fmt.Sprintf("theater in city with id: %d not found.", err.Id)
}
