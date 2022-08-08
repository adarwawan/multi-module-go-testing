package display

import (
	"fmt"

	"github.com/adarwawan/multi-module-go-testing/math"
)

func PrintSum(a int, b int) {
	sum := math.Add(a, b)
	fmt.Println("Sum of %d and %d is %d", a, b, sum)
}
