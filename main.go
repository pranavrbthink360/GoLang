package main

import (
	"fmt"
	"strconv"
)

func main() {
	ans := []string{}
	ans = fizzBuzz(100)

	fmt.Println(ans)
}

func fizzBuzz(n int) []string {
	ans := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans = append(ans, "FizzBuzz")
		} else if i%3 == 0 {
			ans = append(ans, "Fizz")
		} else if i%5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}
