package duration

import "testing"

func TestDurationBetween_5_02_1996_And_10_6_2019_ShouldBe_280_Months_5_Days(t *testing.T) {
	expectedResult := "280 months, 5 days"
	sDay := 27
	sMonth := 10
	sYear := 1997

	eDay := 10
	eMonth := 6
	eYear := 2019

	actualResult := CalulateDayOfOfMonth(sDay, sMonth, sYear, eDay, eMonth, eYear)

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

}
