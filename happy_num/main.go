package main

import "fmt"

func isHappy(num int) bool {
	slow := num
	fast := sumOfSquaredDigits(num)

	for fast != 1 && fast != slow {
		slow = sumOfSquaredDigits(slow)
		fast = sumOfSquaredDigits(sumOfSquaredDigits(fast))
	}

	return fast == 1
}

// pow calculates the power of the given digit
func pow(digit int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}	

// sumOfSquaredDigits is a helper function that calculates the sum of squared digits
func sumOfSquaredDigits(number int) int {
	totalSum := 0
	for number > 0 {
		digit := number % 10
		number = number / 10
		totalSum += pow(digit, 2)
	}
	return totalSum
}

func main() {
	inputs := []int{
		1,
		19,
		8,
		7,
		2147483646,
	}

	for _, i := range inputs {
		fmt.Printf("Input=[%v] , Result=[%v]\n\n", i, isHappy(i))
	}
}
