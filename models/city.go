package models

type CityIdModel int
type CityModel struct {
	Id   CityIdModel `json:"id"`
	Name string      `json:"name"`
}
