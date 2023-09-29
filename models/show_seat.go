package models

import "movie_booking/enums"

type ShowSeatIdModel int
type ShowSeatModel struct {
	Id         ShowSeatIdModel      `json:"id"`
	ShowId     ShowIdModel          `json:"showId"`
	SeatId     SeatIdModel          `json:"seatId"`
	SeatStatus enums.SeatStatusEnum `json:"seatStatus"`
}
