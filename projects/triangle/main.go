package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	q := math.Sqrt(float64(a*a + b*b))
	fmt.Print(q)
}
