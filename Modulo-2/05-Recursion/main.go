package main

import "fmt"

func main() {
	// Recursion
	contar(10)
}

func contar(a int) int {
	if a == 0 {
		return 0
	}
	fmt.Println(a)
	return a + contar(a-1)
}
