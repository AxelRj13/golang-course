package main

import "fmt"

type numbers []int

func newNumbers(n int) numbers {
	nums := numbers{}

	for i := 1; i <= n; i++ {
		nums = append(nums, i) // adding value to nums slice
	}

	return nums
}

func (n numbers) checkEvenOrOdd() {
	for _, num := range n {
		if num%2 == 0 {
			fmt.Printf("%v is Even \n", num)
		} else {
			fmt.Printf("%v is Odd \n", num)
		}
	}
}
