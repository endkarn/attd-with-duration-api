package duration

import (
	"testing"
)
	
func TestConvertDate_16_10_1997_toTimestamp_876960000(t *testing.T){
	expectedResult := 876960000

	actualResult := ConvertDateToTimestamp()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

}