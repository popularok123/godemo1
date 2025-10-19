package main

import "fmt"

type Variafun struct{}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	return total
}

func (vf *Variafun) Do() {
	sum(1, 2, 3)
}
