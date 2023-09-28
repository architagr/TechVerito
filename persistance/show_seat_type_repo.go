package persistance

import (
	"movie_booking/errors"
	"movie_booking/models"
)

type IShowSeatTypePersistance interface {
	GetShowSeatType(showId models.ShowIdModel) (result []models.ShowSeatTypeModel, err error)
}

type ShowSeatTypePersistance struct {
	showSeatTypes []models.ShowSeatTypeModel
}

func (repo *ShowSeatTypePersistance) GetShowSeatType(showId models.ShowIdModel) (result []models.ShowSeatTypeModel, err error) {

	result = make([]models.ShowSeatTypeModel, 0, 10)

	for _, seatType := range repo.showSeatTypes {
		if seatType.ShowId == showId {
			result = append(result, seatType)
		}
	}

	if len(result) == 0 {
		err = &errors.ShowSeatTypeNotFoundError{
			ShowId: showId,
		}
	}
	return
}

func InitShowSeatTypePersistance() IShowSeatTypePersistance {
	return &ShowSeatTypePersistance{
		showSeatTypes: []models.ShowSeatTypeModel{
			{
				Id:         models.ShowSeatTypeIdModel(1),
				ShowId:     models.ShowIdModel(1),
				SeatTypeId: models.SeatTypeIdModel(1),
				Rate:       100,
			},
			{
				Id:         models.ShowSeatTypeIdModel(2),
				ShowId:     models.ShowIdModel(1),
				SeatTypeId: models.SeatTypeIdModel(2),
				Rate:       150,
			},
			{
				Id:         models.ShowSeatTypeIdModel(3),
				ShowId:     models.ShowIdModel(1),
				SeatTypeId: models.SeatTypeIdModel(3),
				Rate:       200,
			},
			{
				Id:         models.ShowSeatTypeIdModel(4),
				ShowId:     models.ShowIdModel(2),
				SeatTypeId: models.SeatTypeIdModel(1),
				Rate:       100,
			},
			{
				Id:         models.ShowSeatTypeIdModel(5),
				ShowId:     models.ShowIdModel(2),
				SeatTypeId: models.SeatTypeIdModel(2),
				Rate:       150,
			},
			{
				Id:         models.ShowSeatTypeIdModel(6),
				ShowId:     models.ShowIdModel(2),
				SeatTypeId: models.SeatTypeIdModel(3),
				Rate:       200,
			},
			{
				Id:         models.ShowSeatTypeIdModel(7),
				ShowId:     models.ShowIdModel(3),
				SeatTypeId: models.SeatTypeIdModel(1),
				Rate:       100,
			},
			{
				Id:         models.ShowSeatTypeIdModel(8),
				ShowId:     models.ShowIdModel(3),
				SeatTypeId: models.SeatTypeIdModel(2),
				Rate:       150,
			},
			{
				Id:         models.ShowSeatTypeIdModel(9),
				ShowId:     models.ShowIdModel(3),
				SeatTypeId: models.SeatTypeIdModel(3),
				Rate:       200,
			},
		},
	}
}
