package main

import (
	"fmt"
	"strings"
)

func A(s string) string {
	var z string = ""
	for i := 0; i < len(s); i++ {
		z += (string(s[i]) + " ")
	}
	return z
}

func main() {
	var a string
	fmt.Scan(&a)
	w := A(a)
	w = strings.Replace(w, " ", "*", -1)
	fmt.Print(string(w[:len(w)-1]))

}
