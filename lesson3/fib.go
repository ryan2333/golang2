package main

import "fmt"

func sum(nums []int) int {
	sums := 0
	for _, num := range nums {
		sums += num
	}
	return sums
}

func main() {
	x, y := 0, 1
	fibs := []int{x, y}
	for y+x < 100 {
		x, y = y, y+x
		fibs = append(fibs, y)
	}
	fmt.Println(fibs)
	fmt.Println(sum(fibs))

}
