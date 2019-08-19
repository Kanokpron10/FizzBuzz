package FizzBuzz

import (
	"testing"
)

func Test_ConvertToFizzBuzz_Input_1_Should_Be_1(t *testing.T) {
	expectedResult := "1"

	actualResult := ConvertToFizzBuzz(1)

	if expectedResult != actualResult {
		t.Errorf("Expect %s but got %s", expectedResult, actualResult)
	}
}

func Test_ConvertToFizzBuzz_Input_3_Should_Be_Fizz(t *testing.T) {
	expectedResult := "Fizz"

	actualResult := ConvertToFizzBuzz(3)

	if expectedResult != actualResult {
		t.Errorf("Expect %s but got %s", expectedResult, actualResult)
	}
}

func Test_ConvertToFizzBuzz_Input_5_Should_Be_Buzz(t *testing.T) {
	expectedResult := "Buzz"

	actualResult := ConvertToFizzBuzz(5)

	if expectedResult != actualResult {
		t.Errorf("Expect %s but got %s", expectedResult, actualResult)
	}
}

func Test_ConvertToFizzBuzz_Input_6_Should_Be_Fizz(t *testing.T) {
	expectedResult := "Fizz"

	actualResult := ConvertToFizzBuzz(6)

	if expectedResult != actualResult {
		t.Errorf("Expect %s but got %s", expectedResult, actualResult)
	}
}

func Test_ConvertToFizzBuzz_Input_15_Should_Be_FizzBuzz(t *testing.T) {
	expectedResult := "FizzBuzz"

	actualResult := ConvertToFizzBuzz(15)

	if expectedResult != actualResult {
		t.Errorf("Expect %s but got %s", expectedResult, actualResult)
	}
}
