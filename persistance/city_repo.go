package persistance

import (
	"movie_booking/errors"
	"movie_booking/models"
)

type ICityPersistance interface {
	GetAllCities() []models.CityModel
	GetCity(id models.CityIdModel) (models.CityModel, error)
	GetCityByName(name string) (models.CityModel, error)
}
type CityPersistance struct {
	cities []models.CityModel
}

func InitCityPersistance() ICityPersistance {
	return &CityPersistance{
		// #region this is inilization of database data if we have db connected then this will not be needed
		cities: []models.CityModel{
			{
				Id:   models.CityIdModel(1),
				Name: "Bangalore",
			},
		},
		// #endregion
	}
}

func (repo *CityPersistance) GetAllCities() []models.CityModel {
	return repo.cities
}
func (repo *CityPersistance) GetCity(id models.CityIdModel) (models.CityModel, error) {

	for _, city := range repo.cities {
		if city.Id == id {
			return city, nil
		}
	}
	return models.CityModel{}, &errors.CityNotFoundError{
		Id: id,
	}
}

func (repo *CityPersistance) GetCityByName(name string) (models.CityModel, error) {

	for _, city := range repo.cities {
		if city.Name == name {
			return city, nil
		}
	}
	return models.CityModel{}, new(errors.CityNotFoundError)
}
