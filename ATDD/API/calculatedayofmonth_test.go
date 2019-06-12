package duration

import "testing"

func TestDurationBetween_5_02_1996_And_10_6_2019_ShouldBe_280_Months_5_Days(t *testing.T) {
	expectedResult := "280 months, 5 days"
	sDay := 5
	sMonth := 2
	sYear := 1996

	eDay := 10
	eMonth := 6
	eYear := 2019

	actualResult := CalculateDayOfOfMonth(sDay, sMonth, sYear, eDay, eMonth, eYear)

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

}

func TestDurationBetween_16_10_1997_And_10_6_2019_ShouldBe_259_Months_25_Days(t *testing.T) {
	expectedResult := "259 months, 25 days"
	sDay := 16
	sMonth := 10
	sYear := 1997

	eDay := 10
	eMonth := 6
	eYear := 2019

	actualResult := CalculateDayOfOfMonth(sDay, sMonth, sYear, eDay, eMonth, eYear)

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

}

func TestDurationBetween_27_10_1997_And_10_6_2019_ShouldBe_259_Months_14_Days(t *testing.T) {
	expectedResult := "259 months, 14 days"
	sDay := 27
	sMonth := 10
	sYear := 1997

	eDay := 10
	eMonth := 6
	eYear := 2019

	actualResult := CalculateDayOfOfMonth(sDay, sMonth, sYear, eDay, eMonth, eYear)

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

}
