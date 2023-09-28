package errors

import (
	"fmt"
	"time"
)

type ShowStartTimeNotFoundError struct {
	StartTime time.Time
}

func (err *ShowStartTimeNotFoundError) Error() string {
	return fmt.Sprintf("shows after %+v are not available", err.StartTime)
}
