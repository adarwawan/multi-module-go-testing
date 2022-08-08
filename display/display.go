package main

import (
	"fmt"

	"github.com/adarwawan/multi-module-go-testing/math"
)

func main() {
	a := 1
	b := 2
	sum := math.Add(a, b)
	fmt.Printf("Sum of %d and %d is %d", a, b, sum)
}
