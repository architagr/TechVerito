package models

type SeatTypeIdModel int
type SeatTypeModel struct {
	Id   SeatTypeIdModel `json:"id"`
	Name string          `json:"name"`
}
