package persistance

import (
	"movie_booking/errors"
	"movie_booking/models"
)

type ISeatTypePersistance interface {
	GetAll() []models.SeatTypeModel
	GetSeatType(id models.SeatTypeIdModel) (seatType models.SeatTypeModel, err error)
}

type SeatTypePersistance struct {
	seatTypes []models.SeatTypeModel
}

func (repo *SeatTypePersistance) GetAll() []models.SeatTypeModel {
	return repo.seatTypes
}
func (repo *SeatTypePersistance) GetSeatType(id models.SeatTypeIdModel) (seatType models.SeatTypeModel, err error) {
	for _, seatTypeData := range repo.seatTypes {
		if seatTypeData.Id == id {
			seatType = seatTypeData
			return
		}
	}

	err = &errors.SeatTypeNotFoundError{
		Id: id,
	}
	return
}

func InitSeatTypePersistance() ISeatTypePersistance {
	return &SeatTypePersistance{
		seatTypes: []models.SeatTypeModel{
			{
				Id:   models.SeatTypeIdModel(1),
				Name: "Silver",
			},
			{
				Id:   models.SeatTypeIdModel(2),
				Name: "Gold",
			},
			{
				Id:   models.SeatTypeIdModel(3),
				Name: "Platinum",
			},
		},
	}
}
