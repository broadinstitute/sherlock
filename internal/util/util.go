package util

import "fmt"

func multBy3(a int) int {
	return a * 3
}

// PrintTriple prints 3 times a to the console
func PrintTriple(a int) {
	fmt.Println(multBy3(a))
}
