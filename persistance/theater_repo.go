package persistance

import (
	"movie_booking/errors"
	"movie_booking/models"
)

type ThearterPersistance struct {
	theaters []models.TheaterModel
}

type IThearterPersistance interface {
	GetAllTheaters() []models.TheaterModel
	GetTheater(id models.TheaterIdModel) (models.TheaterModel, error)
	GetTheaterByCity(id models.CityIdModel) ([]models.TheaterModel, error)
}

func InitThearterPersistance() IThearterPersistance {

	return &ThearterPersistance{
		// #region this is inilization of database data if we have db connected then this will not be needed
		theaters: []models.TheaterModel{
			{
				Id:     models.TheaterIdModel(1),
				Name:   "PVR",
				CityId: models.CityIdModel(1),
			},
			{
				Id:     models.TheaterIdModel(2),
				Name:   "Inox",
				CityId: models.CityIdModel(1),
			},
		},
		// #endregion
	}
}

func (repo *ThearterPersistance) GetAllTheaters() []models.TheaterModel {
	return repo.theaters
}
func (repo *ThearterPersistance) GetTheater(id models.TheaterIdModel) (models.TheaterModel, error) {

	for _, theater := range repo.theaters {
		if theater.Id == id {
			return theater, nil
		}
	}
	return models.TheaterModel{}, &errors.TheaterNotFoundError{
		Id: id,
	}
}
func (repo *ThearterPersistance) GetTheaterByCity(id models.CityIdModel) ([]models.TheaterModel, error) {
	result := make([]models.TheaterModel, 0, len(repo.theaters))

	for _, theater := range repo.theaters {
		if theater.CityId == id {
			result = append(result, theater)
		}
	}

	if len(result) == 0 {
		return result, &errors.TheaterInCityNotFoundError{
			Id: id,
		}
	}

	return result, nil
}
