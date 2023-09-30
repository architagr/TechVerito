package persistance

import (
	"movie_booking/models"
	"testing"
)

func TestGetAllAudi(t *testing.T) {
	audiRepo := InitAudiPersistance()

	t.Run("Should not give error", func(tb *testing.T) {
		_, err := audiRepo.GetAllAudi(models.TheaterIdModel(1))
		if err != nil {
			tb.Error("We should have Audi for theated id 1")
		}
	})
	t.Run("Should give error for theater id that does not exists", func(tb *testing.T) {
		_, err := audiRepo.GetAllAudi(models.TheaterIdModel(10))
		if err == nil {
			tb.Error("We should not have Audi for theated id 10")
		}
	})

}
