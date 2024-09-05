package main

import (
	"errors"
	"fmt"
	"strconv"
)

func convertToBase(number int, base int) (string, error) {

	if base < 2 || base > 36 {
		return "", errors.New("base must be between 2 and 36")
	}
	return strconv.FormatInt(int64(number), base), nil
}

func main() {
	var number, base int

	fmt.Printf("Enter a decimal number: ")
	fmt.Scanf("%d", &number)

	fmt.Printf("Enter the base (between 2 and 36): ")
	fmt.Scanf("%d", &base)

	result, err := convertToBase(number, base)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The number %d in base %d is: %s\n", number, base, result)
	}
}
