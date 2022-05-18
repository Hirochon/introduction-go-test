package main

import "github.com/Hirochon/introduction-go-test/error_handling"

func main() {
	println(sum(2, 2))
	if err := error_handling.RequestValidation(500); err != nil {
		println(err.Error())
	}
}

func sum(a, b int) int {
	return a * b
}
