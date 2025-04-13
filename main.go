package main

import "fmt"

func main() {
	fmt.Println(Add(2, 3))
	fmt.Println((Factorial(2))) // 2
	fmt.Println((Factorial(3))) //6

}

func Add(a, b int) int {
	return a + b
}

func Factorial(n int) (result int) {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	return n * Factorial(n-1)
}

// Factorial = !5 =5*4*3*2*1
