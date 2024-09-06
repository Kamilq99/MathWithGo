package main

import (
	"errors"
	"fmt"
)

func fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Error: n must be a non-negative integer")
	}

	if n == 0 {
		return 0, nil
	}
	if n == 1 {
		return 1, nil
	}

	fib1, err1 := fibonacci(n - 1)
	if err1 != nil {
		return 0, err1
	}

	fib2, err2 := fibonacci(n - 2)
	if err2 != nil {
		return 0, err2
	}

	return fib1 + fib2, nil
}

func main() {
	var iterations int

	fmt.Print("Get number of iterations: ")
	fmt.Scan(&iterations)

	fmt.Println("Fibonacci sequence:")
	for i := 0; i < iterations; i++ {
		result, err := fibonacci(i)
		if err != nil {
			fmt.Printf("Error calculating fibonacci(%d): %s\n", i, err)
			continue
		}
		fmt.Printf("%d ", result)
	}
	fmt.Println()
}
