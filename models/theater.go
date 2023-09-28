package models

type TheaterIdModel int

type TheaterModel struct {
	Id     TheaterIdModel `json:"id"`
	Name   string         `json:"name"`
	CityId CityIdModel    `json:"cityId"`
}
