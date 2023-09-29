package persistance

import (
	"movie_booking/errors"
	"movie_booking/models"
)

type ISeatPersistance interface {
	GetSeats(audiId models.AudiIdModel) (seats []models.SeatModel, err error)
	GetSeatByName(audiId models.AudiIdModel, name string) (seat models.SeatModel, err error)
}

type SeatPersistance struct {
	seats []models.SeatModel
}

func (repo *SeatPersistance) GetSeats(audiId models.AudiIdModel) (seats []models.SeatModel, err error) {
	result := make([]models.SeatModel, 0, len(repo.seats))

	for _, seat := range repo.seats {
		if seat.AudiId == audiId {
			result = append(result, seat)
		}
	}
	if len(result) == 0 {
		err = &errors.SeatsNotFoundError{
			Id: audiId,
		}
	}
	return
}

func (repo *SeatPersistance) GetSeatByName(audiId models.AudiIdModel, name string) (result models.SeatModel, err error) {

	for _, seat := range repo.seats {
		if seat.AudiId == audiId && seat.Name == name {
			result = seat
			break
		}
	}
	err = &errors.SeatsByNameNotFoundError{
		AudiId: audiId,
		Name:   name,
	}
	return
}

func InitSeats() ISeatPersistance {
	return &SeatPersistance{
		seats: []models.SeatModel{
			// #region Audi 1 row 1
			{
				Id:         models.SeatIdModel(1),
				Name:       "A1",
				Row:        0,
				Col:        0,
				SeatTypeId: models.SeatTypeIdModel(1),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(2),
				Name:       "A2",
				Row:        0,
				Col:        1,
				SeatTypeId: models.SeatTypeIdModel(1),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(3),
				Name:       "A3",
				Row:        0,
				Col:        3,
				SeatTypeId: models.SeatTypeIdModel(1),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(4),
				Name:       "A4",
				Row:        0,
				Col:        4,
				SeatTypeId: models.SeatTypeIdModel(1),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(5),
				Name:       "A5",
				Row:        0,
				Col:        5,
				SeatTypeId: models.SeatTypeIdModel(1),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(6),
				Name:       "A6",
				Row:        0,
				Col:        6,
				SeatTypeId: models.SeatTypeIdModel(1),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(7),
				Name:       "A7",
				Row:        0,
				Col:        7,
				SeatTypeId: models.SeatTypeIdModel(1),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(8),
				Name:       "A8",
				Row:        0,
				Col:        8,
				SeatTypeId: models.SeatTypeIdModel(1),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(9),
				Name:       "A9",
				Row:        0,
				Col:        9,
				SeatTypeId: models.SeatTypeIdModel(1),
				AudiId:     models.AudiIdModel(1),
			},
			// #endregion
			// #region Audi 1 row 2
			{
				Id:         models.SeatIdModel(10),
				Name:       "B1",
				Row:        1,
				Col:        0,
				SeatTypeId: models.SeatTypeIdModel(2),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(11),
				Name:       "B2",
				Row:        1,
				Col:        1,
				SeatTypeId: models.SeatTypeIdModel(2),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(12),
				Name:       "B3",
				Row:        1,
				Col:        3,
				SeatTypeId: models.SeatTypeIdModel(2),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(13),
				Name:       "B4",
				Row:        1,
				Col:        4,
				SeatTypeId: models.SeatTypeIdModel(2),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(14),
				Name:       "B5",
				Row:        1,
				Col:        5,
				SeatTypeId: models.SeatTypeIdModel(2),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(15),
				Name:       "B6",
				Row:        1,
				Col:        6,
				SeatTypeId: models.SeatTypeIdModel(2),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(16),
				Name:       "B7",
				Row:        1,
				Col:        7,
				SeatTypeId: models.SeatTypeIdModel(2),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(17),
				Name:       "B8",
				Row:        1,
				Col:        8,
				SeatTypeId: models.SeatTypeIdModel(2),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(18),
				Name:       "B9",
				Row:        1,
				Col:        9,
				SeatTypeId: models.SeatTypeIdModel(2),
				AudiId:     models.AudiIdModel(1),
			},
			// #endregion
			// #region Audi 1 row 3
			{
				Id:         models.SeatIdModel(19),
				Name:       "C1",
				Row:        2,
				Col:        0,
				SeatTypeId: models.SeatTypeIdModel(3),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(20),
				Name:       "C2",
				Row:        2,
				Col:        1,
				SeatTypeId: models.SeatTypeIdModel(3),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(21),
				Name:       "C3",
				Row:        3,
				Col:        3,
				SeatTypeId: models.SeatTypeIdModel(3),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(22),
				Name:       "C4",
				Row:        3,
				Col:        4,
				SeatTypeId: models.SeatTypeIdModel(3),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(23),
				Name:       "C5",
				Row:        2,
				Col:        5,
				SeatTypeId: models.SeatTypeIdModel(3),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(24),
				Name:       "C6",
				Row:        2,
				Col:        6,
				SeatTypeId: models.SeatTypeIdModel(3),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(25),
				Name:       "C7",
				Row:        3,
				Col:        7,
				SeatTypeId: models.SeatTypeIdModel(3),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(26),
				Name:       "C8",
				Row:        3,
				Col:        8,
				SeatTypeId: models.SeatTypeIdModel(3),
				AudiId:     models.AudiIdModel(1),
			},
			{
				Id:         models.SeatIdModel(27),
				Name:       "C9",
				Row:        3,
				Col:        9,
				SeatTypeId: models.SeatTypeIdModel(3),
				AudiId:     models.AudiIdModel(1),
			},
			// #endregion
		},
	}
}
