package main

import (
	"fmt"

	"github.com/xyproto/mathy"
)

func main() {
	fmt.Println("hi")

	// Implement "+"
	plus := mathy.NewOperator("+")
	plus.IntFunc(func(a, b int) int {
		return a + b
	})
	plus.Float64Func(func(a, b float64) float64 {
		return a + b
	})
	plus.StringFunc(func(a, b string) string {
		return a + b
	})

	// Implement "*"
	mul := mathy.NewOperator("*")
	mul.IntFunc(func(a, b int) int {
		return a * b
	})
	mul.Float64Func(func(a, b float64) float64 {
		return a * b
	})
	mul.StringFunc(func(a, b string) string {
		return a + "," + b
	})

	result := mathy.EvalInt("1+2+3+4", plus, mul)

	fmt.Printf("The result is: %d\n", result)

}
