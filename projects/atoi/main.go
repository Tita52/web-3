package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	var result string

	for _, ch := range input {

		digit := int(ch - '0')

		square := digit * digit

		result += fmt.Sprintf("%d", square)
	}

	fmt.Println(result)
}
