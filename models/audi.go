package models

import "movie_booking/enums"

type AudiIdModel int
type AudiModel struct {
	Id        AudiIdModel             `json:"id"`
	Name      string                  `json:"name"`
	TheaterId TheaterIdModel          `json:"theaterId"`
	Features  []enums.AudiFeatureEnum `json:"features"`
}
