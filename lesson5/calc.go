package main

import (
	"fmt"
	"strconv"
)

func calc(args ...string) int {
	sum, _ := strconv.Atoi(args[0])

	for i := 2; i < len(args); i += 2 {
		num, _ := strconv.Atoi(args[i])
		switch args[i-1] {
		case "+":
			sum += num
		case "-":
			sum -= num
		case "*":
			sum *= num
		case "/":
			sum /= num
		default:
			continue
		}
	}

	return sum
}

func main() {
	// fmt.Println(calc("1", "+", "2"))
	fmt.Println(calc("1", "+", "2", "*", "5"))
}
