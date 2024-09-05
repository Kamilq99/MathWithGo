package main

import (
	"errors"
	"fmt"
	"math"
)

func logarithm(a float64, base float64) (float64, error) {
	var result float64

	if base < 2 {
		return 0, errors.New("Error: The base is less than 2")
	} else if a <= 0 {
		return 0, errors.New("Error: The number is less than or equal to 0")
	}

	result = math.Log(a) / math.Log(base)

	return result, nil
}

func main() {

	var a float64
	var base float64
	var result float64

	fmt.Printf("Get number a = ")
	fmt.Scanf("%f", &a)

	fmt.Printf("Get number base = ")
	fmt.Scanf("%f", &base)

	result, err := logarithm(a, base)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result = %f\n", result)
	}
}
