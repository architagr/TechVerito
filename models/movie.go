package models

import "movie_booking/enums"

type MovieIdModel int
type MovieModel struct {
	Id        MovieIdModel         `json:"id"`
	Name      string               `json:"name"`
	Languages []enums.LanguageEnum `json:"languages"`
}
