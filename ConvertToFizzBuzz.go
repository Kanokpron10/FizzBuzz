package FizzBuzz

import "strconv"

func ConvertToFizzBuzz(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz"
	}
	if number%5 == 0 {
		return "Buzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(number)
}
