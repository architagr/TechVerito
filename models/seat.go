package models

type SeatIdModel int
type SeatModel struct {
	Id         SeatIdModel     `json:"id"`
	Name       string          `json:"name"`
	Row        int             `json:"row"`
	Col        int             `json:"col"`
	SeatTypeId SeatTypeIdModel `json:"seatTypeId"`
	AudiId     AudiIdModel     `json:"audiId"`
}
