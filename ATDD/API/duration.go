package duration

import (
	"time"
	"strconv"
)

func ConvertDateToTimestamp(day int, month int, year int) int {
	sec := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return int(sec.Unix())
	
}

func ConvertSecondsToDays(second int) string {
	days := strconv.Itoa(second/86400)
	return days +" days"
}

func Duration(start int,end int) int {
	diff := end-start
	return diff 
}

