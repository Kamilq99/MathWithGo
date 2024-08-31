package main

import (
	"errors"
	"fmt"
)

func checkIfYouCanBuildTraingle(a int, b int, c int) (bool, error) {

	var result bool

	if a+b > c || a+c > b || b+c > a {
		return result, errors.New("Error: You can't build a triangle")
	}

	return result, nil

}
func main() {

	var a int
	var b int
	var c int
	var result bool
	var err error

	fmt.Printf("Get side a = ")
	fmt.Scanf("%d", &a)

	fmt.Printf("Get side b = ")
	fmt.Scanf("%d", &b)

	fmt.Printf("Get side c = ")
	fmt.Scanf("%d", &c)

	result, err = checkIfYouCanBuildTraingle(a, b, c)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result = %t\n", result)
	}
}
