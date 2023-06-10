package go_testing

func Add(numbers ...int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return sum
}
