package models

import "movie_booking/enums"

type ShowResponse struct {
	ShowId      ShowIdModel             `json:"showId"`
	MovieName   string                  `json:"movieName"`
	Language    enums.LanguageEnum      `json:"language"`
	Features    []enums.AudiFeatureEnum `json:"features"`
	TheaterName string                  `json:"theaterName"`
	AudiName    string                  `json:"AudiName"`
	Seats       [][]SeatInfo            `json:"seats"`
}
