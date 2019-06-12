package duration

import (
	"testing"
)

func TestConvertSecond_683164800_toDay_7907(t *testing.T){
	expectedResult := "7907 days"
	second := 683164800

	actualResult := ConvertSecondsToDays(second)

	if actualResult != expectedResult{
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestConvertSecond_136646400_toDay_8526(t *testing.T){
	expectedResult := "8526 days"
	second := 736646400

	actualResult := ConvertSecondsToDays(second)

	if actualResult != expectedResult{
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestConvertSecond_682214400_toDay_7896(t *testing.T){
	expectedResult := "7896 days"
	second := 682214400

	actualResult := ConvertSecondsToDays(second)

	if actualResult != expectedResult{
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}