package main

import (
	"fmt"

	"github.com/YongHaoWu/betterGo/enum"
)

func mul(a, b int) int {
	return a * b
}

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}
	// Compute 10!
	out := enum.Reduce(a, mul, 1).(int)
	expect := 1
	for i := range a {
		expect *= a[i]
	}
	if expect != out {
		fmt.Printf("expected %d got %d", expect, out)
	}
	fmt.Println("success, ", expect)
}
