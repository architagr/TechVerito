package errors

import (
	"fmt"
	"movie_booking/enums"
)

type MovieInLanguageNotFoundError struct {
	Lang enums.LanguageEnum
}

func (err *MovieInLanguageNotFoundError) Error() string {
	return fmt.Sprintf("movie in language: %s not found.", err.Lang.ToString())
}
