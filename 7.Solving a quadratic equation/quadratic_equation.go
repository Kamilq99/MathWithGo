package main

import (
	"errors"
	"fmt"
)

func QudraticEquation(a int, b int, c int) (int, error) {

	var delta int

	delta = b*b - 4*a*c

	if delta < 0 {
		return 0, errors.New("Error: delta is less than 0")
	} else if delta == 0 {
		return -b / (2 * a), nil
	} else {
		return (-b + delta) / (2 * a), nil
	}
}

func main() {

	var a int
	var b int
	var c int

	fmt.Printf("Get a = ")
	fmt.Scanf("%d", &a)
	fmt.Scanln()
	fmt.Printf("Get b = ")
	fmt.Scanf("%d", &b)
	fmt.Scanln()
	fmt.Printf("Get c = ")
	fmt.Scanf("%d", &c)
	fmt.Scanln()

	result, err := QudraticEquation(a, b, c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result = %d\n", result)
	}

}
