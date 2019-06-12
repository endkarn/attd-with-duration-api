package duration

import (
	"fmt"
	"strconv"
	"time"
)

func ConvertDateToTimestamp(day int, month int, year int) int {
	sec := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return int(sec.Unix())

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

func Duration(start int, end int) int {
	diff := end - start
	return diff
}
