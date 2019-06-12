package duration

import (
	"testing"
)
	
func TestConvertDate_16_10_1997_ShouldGetTimestamp_876960000(t *testing.T){
	expectedResult := 876960000
	day := 16
	month := 10
	year := 1997

	actualResult := ConvertDateToTimestamp(day,month,year)

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

}

func TestConvertDate_27_10_1997_ShouldGetTimestamp_877910400(t *testing.T){
	expectedResult := 877910400
	day := 27
	month := 10
	year := 1997

	actualResult := ConvertDateToTimestamp(day,month,year)

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

}

