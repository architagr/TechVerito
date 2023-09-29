package models

import "movie_booking/enums"

type SeatInfo struct {
	Name         string               `json:"seatName"`
	SeatTypeName string               `json:"seatTypeName"`
	SeatId       SeatIdModel          `json:"-"`
	Rate         float32              `json:"rate"`
	Status       enums.SeatStatusEnum `json:"status"`
}
type ShowSeatResponse struct {
	Seats [][]SeatInfo `json:"seats"`
}
