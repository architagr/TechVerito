package models

import "movie_booking/enums"

type ShowSeatModel struct {
	Id         int                  `json:"id"`
	ShowId     ShowIdModel          `json:"showId"`
	SeatId     SeatIdModel          `json:"seatId"`
	SeatStatus enums.SeatStatusEnum `json:"seatStatus"`
}
