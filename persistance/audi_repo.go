package persistance

import (
	"movie_booking/enums"
	"movie_booking/errors"
	"movie_booking/models"
)

type IAudiPersistance interface {
	GetAllAudi(theaterId models.TheaterIdModel) ([]models.AudiModel, error)
	GetAudi(audiId models.AudiIdModel) (models.AudiModel, error)
}

type AudiPersistance struct {
	audies []models.AudiModel
}

func InitAudiPersistance() IAudiPersistance {
	return &AudiPersistance{
		// #region this is inilization of database data if we have db connected then this will not be needed
		audies: []models.AudiModel{
			{
				Id:        models.AudiIdModel(1),
				Name:      "Audi 1",
				TheaterId: models.TheaterIdModel(1),
				Features: []enums.AudiFeatureEnum{
					enums.AUDI_FEATURE_THREE_DIMINTIONAL,
					enums.AUDI_FEATURE_TWO_DIMINTIONAL,
				},
			},
			{
				Id:        models.AudiIdModel(2),
				Name:      "Audi 2",
				TheaterId: models.TheaterIdModel(1),
				Features: []enums.AudiFeatureEnum{
					enums.AUDI_FEATURE_THREE_DIMINTIONAL,
					enums.AUDI_FEATURE_TWO_DIMINTIONAL,
					enums.AUDI_FEATURE_DOLBY,
				},
			},
			{
				Id:        models.AudiIdModel(3),
				Name:      "Audi 1",
				TheaterId: models.TheaterIdModel(2),
				Features: []enums.AudiFeatureEnum{
					enums.AUDI_FEATURE_THREE_DIMINTIONAL,
					enums.AUDI_FEATURE_TWO_DIMINTIONAL,
					enums.AUDI_FEATURE_IMAX,
				},
			},
			{
				Id:        models.AudiIdModel(4),
				Name:      "Audi 2",
				TheaterId: models.TheaterIdModel(2),
				Features: []enums.AudiFeatureEnum{
					enums.AUDI_FEATURE_TWO_DIMINTIONAL,
					enums.AUDI_FEATURE_DOLBY,
				},
			},
		},
		// #endregion
	}
}

func (repo *AudiPersistance) GetAllAudi(theaterId models.TheaterIdModel) (result []models.AudiModel, err error) {
	result = make([]models.AudiModel, 0, len(repo.audies))

	for _, audi := range repo.audies {
		if audi.TheaterId == theaterId {
			result = append(result, audi)
		}
	}
	if len(result) == 0 {
		err = &errors.AudiInTheaterNotFoundError{
			Id: theaterId,
		}
	}
	return
}
func (repo *AudiPersistance) GetAudi(audiId models.AudiIdModel) (models.AudiModel, error) {
	for _, audi := range repo.audies {
		if audi.Id == audiId {
			return audi, nil
		}
	}
	return models.AudiModel{}, &errors.AudiNotFoundError{
		Id: audiId,
	}
}
