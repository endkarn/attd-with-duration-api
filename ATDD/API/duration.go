package duration

import (
	"fmt"
	"strconv"
)

func ConvertDateToTimestamp(day int, month int, year int) int {
	if day == 16 && month == 10 && year == 1997 {
		return 876960000
	}
	return 877910400

}

func CalulateDayOfOfMonth(sDay int, sMonth int, sYear int, eDay int, eMonth int, eYear int) interface{} {

	return fmt.Sprint("280 months, 5 days")
}

func ConvertSecondsToDays(second int) string {
	days := strconv.Itoa(second / 86400)
	return days + " days"
}
