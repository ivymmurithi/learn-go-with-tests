package sumoflsit

func GenerateSum(numbers[4]int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func GenerateDynamicListSum(numbers[]int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		//append is flexible and can expand the underlying array size
		// when you know the numbers of items you use make
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// below a new slice 'tail' is created with copies of the value 
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
		
	}
	return sums	
}