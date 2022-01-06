package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	var name string
	var age uint
	var score uint = 0
	var questions uint = 2
	var cores uint

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay, you can play!")
	} else {
		fmt.Println("Sorry, you can't play!")
		return
	}

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer1, answer2 string
	fmt.Scan(&answer1, &answer2)

	var answer string = answer1 + " " + answer2

	if answer == "RTX 3090" || answer == "rtx 3090" {
		score++
		fmt.Println("Correct.")
	} else {
		fmt.Println("Incorrect.")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	fmt.Scan(&cores)

	if cores == 12 {
		score++
		fmt.Println("Correct.")
	} else {
		fmt.Print("Inncorrect.")
	}

	var percent = (float64(score) / float64(questions)) * 100

	fmt.Printf("Your score is: %v out of %v.\n", score, questions)
	fmt.Printf("Your scored: %v%%\n", percent)
}
