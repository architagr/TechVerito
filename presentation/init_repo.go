package presentation

import (
	"movie_booking/persistance"
)

var audiRepo persistance.IAudiPersistance
var cityRepo persistance.ICityPersistance
var movieRepo persistance.IMoviePersistance
var seatRepo persistance.ISeatPersistance
var seatTypeRepo persistance.ISeatTypePersistance
var showRepo persistance.IShowPersistance
var showSeatRepo persistance.IShowSeatPersistance
var showSeatTypeRepo persistance.IShowSeatTypePersistance
var theaterRepo persistance.IThearterPersistance
var ticketRepo persistance.ITicketPersistance

func InitAllPersistance() {
	audiRepo = persistance.InitAudiPersistance()
	cityRepo = persistance.InitCityPersistance()
	movieRepo = persistance.InitMoviePersistance()
	seatRepo = persistance.InitSeats()
	seatTypeRepo = persistance.InitSeatTypePersistance()
	showRepo = persistance.InitShow()
	showSeatRepo = persistance.InitShowSeatPersistance()
	showSeatTypeRepo = persistance.InitShowSeatTypePersistance()
	theaterRepo = persistance.InitThearterPersistance()
	ticketRepo = persistance.InitTicketPersistance()
}
