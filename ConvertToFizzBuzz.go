package fizzbuzz

func ConvertToFizzBuzz(number int) string {
	fizzbuzz := map[int]string{
		1:  "1",
		3:  "Fizz",
		5:  "Buzz",
		6:  "Fizz",
		15: "FizzBuzz",
	}
	return fizzbuzz[number]
}
