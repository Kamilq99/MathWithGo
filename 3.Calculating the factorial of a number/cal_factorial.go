package main

import (
	"errors"
	"fmt"
)

func calculateFactorial(a int) (int, error) {

	if a < 0 {
		return 0, errors.New("Error: The number is negative")
	} else if a == 0 {
		return 1, nil
	}

	for i := a - 1; i > 0; i-- {
		a *= i

		if a == 0 {
			break
		}
	}
	return a, nil
}

func main() {

	var a int
	var result int

	fmt.Printf("Get number a = ")
	fmt.Scanf("%d", &a)

	result, err := calculateFactorial(a)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result = %d\n", result)
	}
}
