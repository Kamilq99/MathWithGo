package main

import (
	"errors"
	"fmt"
)

func checkPrime(a int) (bool, error) {

	if a <= 1 {
		return false, errors.New("Error: The number 1 is not prime")
	}

	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false, errors.New("Error: The number is not prime")
		}
	}

	return true, nil
}

func main() {

	var a int
	var result bool

	fmt.Printf("Get number a = ")
	fmt.Scanf("%d", &a)

	result, err := checkPrime(a)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result = %t\n", result)
	}
}
