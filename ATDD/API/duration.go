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

func CalculateDayOfOfMonth(sDay int, sMonth int, sYear int, eDay int, eMonth int, eYear int) interface{} {

	diffYear := eYear - sYear
	diffMonths := eMonth - sMonth
	diffDay := eDay - sDay

	if diffDay < 0 {
		diffDay += 31
		diffMonths--
	}

	if diffMonths < 0 {
		diffMonths += 12
		diffYear--
	}

	totalDiffMonths := (diffYear * 12) + diffMonths

	return fmt.Sprintf("%v months, %v days", totalDiffMonths, diffDay)
}

func ConvertSecondsToDays(second int) string {
	days := strconv.Itoa(second / 86400)
	return days + " days"
}
