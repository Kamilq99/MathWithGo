package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {

	var result int

	if b == 0 {
		return 0, errors.New("Error: Divide by zero")
	}

	result = a / b

	return result, nil
}

func main() {

	var a int
	var b int
	var result int
	var err error

	fmt.Printf("Get numer a = ")
	fmt.Scanf("%d", &a)

	fmt.Printf("Get numer b = ")
	fmt.Scanf("%d", &b)

	result, err = divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result = %d\n", result)
	}

}
