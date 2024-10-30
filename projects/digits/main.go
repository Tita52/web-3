package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	var max byte = '0'

	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	fmt.Printf("%c\n", max)

}
