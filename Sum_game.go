package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum_nums() {
	min := 0
	max := 100

	// set seed
	rand.Seed(time.Now().UnixNano())

	//generate random numbers and print question on console

	num0 := rand.Intn(max-min) + min
	num1 := rand.Intn(max-min) + min

	// Create the correct answer to the mathematical question

	Ans := num0 + num1

	// A variable of type int for the users answer

	var ans int

	// Ask the question

	fmt.Println("What is ", num0, " + ", num1, "?")

	// Scan in the users answer

	fmt.Scan(&ans)

	// Compare the users answer to the correct answer

	fmt.Println("You said ", ans, " the correct answer is ", Ans)

}

func main() {
	for i := 1; i <= 4; i++ {
		sum_nums()
	}
}
