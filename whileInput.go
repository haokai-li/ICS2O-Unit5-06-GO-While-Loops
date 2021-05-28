// Created by: haokai
// Created on: May 2021
// This program displays, "While-Loops"

package main

import (
	"fmt"
)

func main() {

	// This function does addition
	var first int
	var second int
	var number int
	var time int
	// input
	fmt.Println("This program is multiple.")
	fmt.Println()
	fmt.Print("Enter first number: ")
	fmt.Scanln(&first)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&second)

	// detect
	for time < second {
		number = number + first
		time ++
	}
	fmt.Println("Answer:", number)
	fmt.Println("\nDone.")
}