package persistance

import (
	"movie_booking/enums"
	"movie_booking/errors"
	"movie_booking/models"
	"time"
)

type IShowPersistance interface {
	GetAllShows(startDate time.Time) (result []models.ShowModel, err error)
	GetShow(id models.ShowIdModel) (models.ShowModel, error)
}

type ShowPersistance struct {
	shows []models.ShowModel
}

func (repo *ShowPersistance) GetAllShows(startDate time.Time) (result []models.ShowModel, err error) {
	result = make([]models.ShowModel, 0, len(repo.shows))

	for _, show := range repo.shows {
		// todo: improve this logic to have date greated then startDate
		if show.StartTime.Day() == startDate.Day() && show.StartTime.Month() == startDate.Month() && show.StartTime.Year() == startDate.Year() {
			result = append(result, show)
		}
	}

	if len(result) == 0 {
		err = &errors.ShowStartTimeNotFoundError{
			StartTime: startDate,
		}
	}
	return
}
func (repo *ShowPersistance) GetShow(id models.ShowIdModel) (models.ShowModel, error) {

	for _, show := range repo.shows {
		if show.Id == id {
			return show, nil
		}
	}
	return models.ShowModel{}, &errors.ShowNotFoundError{
		Id: id,
	}
}

func InitShow() IShowPersistance {
	now := time.Now()

	return &ShowPersistance{
		// #region data
		shows: []models.ShowModel{
			{
				Id:        models.ShowIdModel(1),
				AudiId:    models.AudiIdModel(1),
				StartTime: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, nil),
				EndTime:   time.Date(now.Year(), now.Month(), now.Day(), 3, 0, 0, 0, nil),
				MovieId:   models.MovieIdModel(1),
				Language:  enums.LANGUAGE_HINDI,
				Features: []enums.AudiFeatureEnum{
					enums.AUDI_FEATURE_TWO_DIMINTIONAL,
				},
			},
			{
				Id:        models.ShowIdModel(2),
				AudiId:    models.AudiIdModel(1),
				StartTime: time.Date(now.Year(), now.Month(), now.Day(), 3, 0, 0, 0, nil),
				EndTime:   time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, nil),
				MovieId:   models.MovieIdModel(1),
				Language:  enums.LANGUAGE_HINDI,
				Features: []enums.AudiFeatureEnum{
					enums.AUDI_FEATURE_TWO_DIMINTIONAL,
				},
			},
			{
				Id:        models.ShowIdModel(3),
				AudiId:    models.AudiIdModel(1),
				StartTime: time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, nil),
				EndTime:   time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, nil),
				MovieId:   models.MovieIdModel(1),
				Language:  enums.LANGUAGE_HINDI,
				Features: []enums.AudiFeatureEnum{
					enums.AUDI_FEATURE_TWO_DIMINTIONAL,
				},
			},
		},
		// #endregion
	}
}
