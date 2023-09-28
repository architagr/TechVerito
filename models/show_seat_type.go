package models

type ShowSeatTypeIdModel int
type ShowSeatTypeModel struct {
	Id         ShowSeatTypeIdModel `json:"id"`
	ShowId     ShowIdModel         `json:"showId"`
	SeatTypeId SeatTypeIdModel     `json:"seatTypeId"`
	Rate       float32             `json:"rate"`
}
