package service

import (
	"fmt"
	"movie_booking/models"
	"movie_booking/persistance"
	"time"
)

type IShowService interface {
	GetAllShow() (result []models.ShowResponse, err error)
}

func InitShowService(showRepo persistance.IShowPersistance,
	theaterRepo persistance.IThearterPersistance,
	audiRepo persistance.IAudiPersistance,
	movieRepo persistance.IMoviePersistance,
	showSeatService IShowSeatService) IShowService {
	return &ShowService{
		showRepo:        showRepo,
		theaterRepo:     theaterRepo,
		audiRepo:        audiRepo,
		movieRepo:       movieRepo,
		showSeatService: showSeatService,
	}

}

type ShowService struct {
	showRepo        persistance.IShowPersistance
	theaterRepo     persistance.IThearterPersistance
	audiRepo        persistance.IAudiPersistance
	movieRepo       persistance.IMoviePersistance
	showSeatService IShowSeatService
}

func (svc *ShowService) GetAllShow() (result []models.ShowResponse, err error) {
	result = make([]models.ShowResponse, 0)
	allShows, err := svc.showRepo.GetAllShows(time.Now())
	if err != nil {
		return
	}

	for _, show := range allShows {
		audi, err1 := svc.audiRepo.GetAudi(show.AudiId)
		if err1 != nil {
			err = err1
			return
		}
		theater, err1 := svc.theaterRepo.GetTheater(audi.TheaterId)
		if err1 != nil {
			err = err1
			return
		}
		movie, err1 := svc.movieRepo.GetMovie(show.MovieId)
		if err1 != nil {
			err = err1
			return
		}
		allSeats, err1 := svc.showSeatService.GetAvailableSeats(show.Id)
		if err1 != nil {
			err = err1
			return
		}
		result = append(result, models.ShowResponse{
			ShowId:      show.Id,
			MovieName:   movie.Name,
			Language:    show.Language,
			Features:    show.Features,
			TheaterName: theater.Name,
			AudiName:    audi.Name,
			Seats:       allSeats,
		})
	}
	if len(result) == 0 {
		err = fmt.Errorf("no shows at the moment")
	}
	return
}
