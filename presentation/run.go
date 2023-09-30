package presentation

import (
	"bufio"
	"fmt"
	"movie_booking/enums"
	"movie_booking/models"
	"os"
	"strconv"
	"strings"
)

var in *bufio.Reader

func ReadString(in *bufio.Reader) string {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	nStr = strings.TrimSpace(nStr)
	nStr = strings.Trim(nStr, "\t \n")
	return nStr
}
func Run() {
	in = bufio.NewReader(os.Stdin)
	initAllPersistance()
	initAllServices()
	showAllShows()
	for {
		fmt.Printf("\nEnter Show no: ")

		showId, _ := strconv.Atoi(ReadString(in))
		seats, err := showSeatService.GetAvailableSeats(models.ShowIdModel(showId))
		if err != nil {
			panic(err)
		}
		fmt.Printf("Available Seats:\n")
		PrintAvailableSeats(seats)
		bookSeats(models.ShowIdModel(showId))
		fmt.Printf("Would you like to continue (Yes/No): ")
		continueInput := ReadString(in)
		if strings.EqualFold(continueInput, "no") {
			break
		}

	}
	printSalesReport()
}
func bookSeats(showId models.ShowIdModel) {
	fmt.Printf("Enter Seats: ")
	allSeats := ReadString(in)
	seatNames := getSaniatisedSeatName(allSeats)
	ticket, err := ticketService.BookTicket(showId, seatNames)
	if err != nil {
		fmt.Printf("%s, Please select different seats\n", err.Error())
	} else {

		fmt.Printf("%s\n", ticket.ToString())
	}
}
func getSaniatisedSeatName(allSeats string) []string {
	splitArr := strings.Split(allSeats, ",")
	result := make([]string, 0, len(splitArr))
	for _, seat := range splitArr {
		result = append(result, strings.Trim(seat, " "))
	}
	return result
}
func printSalesReport() {
	report := salesService.GetTotalSalesReport()
	fmt.Printf("Total Sales:\n")
	fmt.Printf("Revenue: Rs. %.2f\n", report.Revenue)
	for name, value := range report.ColectedTaxes {
		fmt.Printf("%s: Rs. %.2f\n", name, value)
	}
}
func showAllShows() {
	shows, err := showService.GetAllShow()
	if err != nil {
		panic(err)
	}
	for _, show := range shows {
		fmt.Printf("Show %d running in %s\n", show.ShowId, show.AudiName)
		fmt.Printf("All Seats:\n")
		PrintAvailableSeats(show.Seats)
		fmt.Println()
	}

	fmt.Println("-----")
}

func PrintAvailableSeats(seats [][]models.SeatInfo) {
	for _, row := range seats {
		for _, seat := range row {
			val := "__"
			if seat.Status == enums.SEAT_STATUS_AVAILABLE {
				val = seat.Name
			}
			fmt.Printf("%s ", val)
		}
		fmt.Printf("\n")
	}
}
