package duration

import (
	"testing"
)

func TestDurationBetween_Start_876960000_AndEnd_1560124800_ShouldGet_683164800(t *testing.T) {
	expectedResult := 683164800

	actualResult := Duration(876960000, 1560124800)

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestDurationBetween_Start_877910400_AndEnd_1560124800_ShouldGet_682214400(t *testing.T) {
	expectedResult := 682214400

	actualResult := Duration(877910400, 1560124800)

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}
