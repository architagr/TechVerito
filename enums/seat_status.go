package enums

type SeatStatusEnum int

const (
	SEAT_STATUS_AVAILABLE     SeatStatusEnum = iota
	SEAT_STATUS_LOCKED        SeatStatusEnum = iota
	SEAT_STATUS_BOOKED        SeatStatusEnum = iota
	SEAT_STATUS_NOT_AVAILABLE SeatStatusEnum = iota
)
