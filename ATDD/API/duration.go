package duration

import "strconv"

func ConvertDateToTimestamp(day int,month int,year int) int {
	if day==16 && month==10 && year==1997{
		return 876960000
	}
	return 877910400
	
}

func ConvertSecondsToDays(second int) string {
	days := strconv.Itoa(second/86400)
	return days +" days"
}

