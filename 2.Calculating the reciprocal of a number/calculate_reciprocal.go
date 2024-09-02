package main

import (
	"errors"
	"fmt"
)

func calculateReciprocal(a float64) (float64, error) {

	if a == 0 {
		return 0, errors.New("You can't use 0 as a number")
	}

	return 1 / a, nil
}

func main() {
	var a float64
	var reciprocal float64

	fmt.Printf("Get number a = ")
	fmt.Scanf("%f", &a)

	reciprocal, err := calculateReciprocal(a)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Reciprocal = %f\n", reciprocal)
	}
}
