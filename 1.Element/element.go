package main

import (
	"errors"
	"fmt"
	"math"
)

func elementANumber(a float64) (float64, error) {

	var element float64

	if a < 0 {
		return 0, errors.New("Error: The number is negative")
	}

	element = math.Sqrt(a)

	return element, nil
}

func main() {

	var a float64
	var err error
	var element float64

	fmt.Printf("Get number a = ")
	fmt.Scanf("%f", &a)

	element, err = elementANumber(a)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result = %f\n", element)
	}
}
