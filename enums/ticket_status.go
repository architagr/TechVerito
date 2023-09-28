package enums

type TicketStatusEnum int

const (
	TICKET_STATUS_CANCELLED TicketStatusEnum = iota
	TICKET_STATUS_BOOKED    TicketStatusEnum = iota
)
