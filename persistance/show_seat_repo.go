package persistance

import (
	"movie_booking/enums"
	"movie_booking/errors"
	"movie_booking/models"
)

type IShowSeatPersistance interface {
	GetShowSeat(showId models.ShowIdModel) (result []models.ShowSeatModel, err error)
	UpdateShowSeatStatus(showId models.ShowIdModel, seatId models.SeatIdModel, status enums.SeatStatusEnum) error
}

type ShowSeatPersistance struct {
	showSeats []models.ShowSeatModel
}

func (repo *ShowSeatPersistance) GetShowSeat(showId models.ShowIdModel) (result []models.ShowSeatModel, err error) {
	result = make([]models.ShowSeatModel, 0, len(repo.showSeats))

	for _, showSeat := range repo.showSeats {
		if showSeat.ShowId == showId {
			result = append(result, showSeat)
		}
	}
	if len(result) == 0 {
		err = &errors.ShowSeatsNotFoundError{
			ShowId: showId,
		}
	}
	return
}
func (repo *ShowSeatPersistance) UpdateShowSeatStatus(showId models.ShowIdModel, seatId models.SeatIdModel, status enums.SeatStatusEnum) error {

	seatFound := false
	for _, showSeat := range repo.showSeats {
		if showSeat.ShowId == showId && showSeat.SeatId == seatId {
			showSeat.SeatStatus = status
			seatFound = true
		}
	}
	if !seatFound {
		return &errors.ShowSeatNotFoundError{
			ShowId: showId,
			SeatId: seatId,
		}
	}
	return nil
}

func InitShowSeatPersistance() IShowSeatPersistance {
	data := make([]models.ShowSeatModel, 0)
	i := 1
	for showId := 1; showId <= 3; showId++ {
		for seatId := 1; seatId <= 27; seatId++ {
			data = append(data, models.ShowSeatModel{

				Id:         models.ShowSeatIdModel(i),
				ShowId:     models.ShowIdModel(showId),
				SeatId:     models.SeatIdModel(seatId),
				SeatStatus: enums.SEAT_STATUS_AVAILABLE,
			})
			i++
		}
	}
	return &ShowSeatPersistance{
		showSeats: data,
	}
}
