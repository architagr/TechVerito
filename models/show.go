package models

import (
	"movie_booking/enums"
	"time"
)

type ShowIdModel int
type ShowModel struct {
	Id        ShowIdModel             `json:"id"`
	AudiId    AudiIdModel             `json:"audiId"`
	StartTime time.Time               `json:"startTime"`
	EndTime   time.Time               `json:"endTime"`
	MovieId   MovieIdModel            `json:"movieId"`
	Language  enums.LanguageEnum      `json:"language"`
	Features  []enums.AudiFeatureEnum `json:"features"`
}
