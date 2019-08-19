package fizzbuzz

func ConvertToFizzBuzz(number int) string {
	fizzbuzz := map[int]string{
		1: "1",
	}

	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz"
	}
	if number%5 == 0 {
		return "Buzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	return fizzbuzz[number]
}
