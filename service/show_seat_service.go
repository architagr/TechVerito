package service

import (
	"fmt"
	"movie_booking/constants"
	"movie_booking/enums"
	"movie_booking/models"
	"movie_booking/persistance"
)

type SeatPriceDetails struct {
	name string
	rate float32
}
type IShowSeatService interface {
	GetAvailableSeats(showId models.ShowIdModel) ([][]models.SeatInfo, error)
	CheckAvailability(showId models.ShowIdModel, seatName []string) ([]models.SeatInfo, error)
}

type showSeatService struct {
	showRepo         persistance.IShowPersistance
	showSeatRepo     persistance.IShowSeatPersistance
	seatTypeRepo     persistance.ISeatTypePersistance
	showSeatTypeRepo persistance.IShowSeatTypePersistance
	seatRepo         persistance.ISeatPersistance
}

func InitShowSeatService(showRepo persistance.IShowPersistance,
	showSeatRepo persistance.IShowSeatPersistance,
	seatTypeRepo persistance.ISeatTypePersistance,
	showSeatTypeRepo persistance.IShowSeatTypePersistance,
	seatRepo persistance.ISeatPersistance) IShowSeatService {
	return &showSeatService{
		showRepo:         showRepo,
		showSeatRepo:     showSeatRepo,
		seatTypeRepo:     seatTypeRepo,
		showSeatTypeRepo: showSeatTypeRepo,
		seatRepo:         seatRepo,
	}
}
func (svc *showSeatService) GetAvailableSeats(showId models.ShowIdModel) ([][]models.SeatInfo, error) {

	show, err := svc.showRepo.GetShow(showId)
	if err != nil {
		return nil, err
	}
	showSeatStatus, err := svc.getShowSeats(showId)
	if err != nil {
		return nil, err
	}

	showSeatsTypes, err := svc.showSeatTypeRepo.GetShowSeatType(showId)
	if err != nil {
		return nil, err
	}

	audiSeats, err := svc.seatRepo.GetSeats(show.AudiId)
	if err != nil {
		return nil, err
	}

	seatTypes := make(map[models.SeatTypeIdModel]SeatPriceDetails)
	for _, showSeatType := range showSeatsTypes {
		seatsType, err := svc.seatTypeRepo.GetSeatType(showSeatType.SeatTypeId)
		if err != nil {
			return nil, err
		}
		seatTypes[showSeatType.SeatTypeId] = SeatPriceDetails{
			name: seatsType.Name,
			rate: showSeatType.Rate,
		}

	}
	seatStatusResponse := svc.getBlankLayout()
	for _, seat := range audiSeats {

		seatInfo := seatTypes[seat.SeatTypeId]
		seatData := models.SeatInfo{
			Name:         seat.Name,
			SeatTypeName: seatInfo.name,
			SeatId:       seat.Id,
			Rate:         seatInfo.rate,
			Status:       showSeatStatus[seat.Id],
		}
		seatStatusResponse[seat.Row][seat.Col] = seatData
	}
	return seatStatusResponse, nil
}

func (svc *showSeatService) CheckAvailability(showId models.ShowIdModel, seatNames []string) (result []models.SeatInfo, err error) {
	errorSeat := make([]string, 0, len(seatNames))
	seatNameMap := make(map[string]bool)
	for _, name := range seatNames {
		seatNameMap[name] = true
	}
	availableSeats, err := svc.GetAvailableSeats(showId)
	if err != nil {
		return
	}
	result = make([]models.SeatInfo, 0, len(seatNames))

	for _, row := range availableSeats {
		for _, seat := range row {
			if seatNameMap[seat.Name] && seat.Status != enums.SEAT_STATUS_AVAILABLE {
				errorSeat = append(errorSeat, seat.Name)
			} else {
				result = append(result, seat)
			}
		}
	}

	if len(errorSeat) == 0 {
		return
	}
	err = fmt.Errorf("%+v is/are not available", errorSeat)
	return
}

func (svc *showSeatService) getShowSeats(showId models.ShowIdModel) (map[models.SeatIdModel]enums.SeatStatusEnum, error) {
	showSeats, err := svc.showSeatRepo.GetShowSeat(showId)
	if err != nil {
		return nil, err
	}
	showSeatStatus := make(map[models.SeatIdModel]enums.SeatStatusEnum)
	for _, showSeat := range showSeats {
		showSeatStatus[showSeat.SeatId] = showSeat.SeatStatus
	}
	return showSeatStatus, nil
}
func (svc *showSeatService) getBlankLayout() [][]models.SeatInfo {
	seatStatusResponse := make([][]models.SeatInfo, constants.SEAT_LAYOUT_MAX_SIZE)
	for i := 0; i < constants.SEAT_LAYOUT_MAX_SIZE; i++ {
		seatStatusResponse[i] = make([]models.SeatInfo, constants.SEAT_LAYOUT_MAX_SIZE)
	}
	return seatStatusResponse
}
