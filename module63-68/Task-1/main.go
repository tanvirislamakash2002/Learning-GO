package main

import "fmt"

func main() {
	displayMenu := func() {
		fmt.Println("Welcome to grade calculator")
		fmt.Println("1) Calculate grade")
		fmt.Println("2) Check pass/fail status")
		fmt.Println("3) Exit the program")
		fmt.Print("Choose an option: ")
	}

	var choice int
	var score int
	running := true

	for running {
		displayMenu()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter a score (0-100): ")
			fmt.Scan(&score)
			result := calculateGrade(score)

			if result == "Invalid score" {
				fmt.Println("Choose a valid number from 0 - 100")
			}
			fmt.Println("You got", result)
		case 2:
			fmt.Println("check pass fail")
		case 3:
			fmt.Println("exiting program ...")
			running = false
		default:
			fmt.Println("Choose a valid option from the menu")
		}
	}
}

func calculateGrade(score int) string {
	if score >= 90 && score <= 100 {
		return "A"
	} else if score >= 80 && score <= 89 {
		return "B"
	} else if score >= 70 && score <= 79 {
		return "C"
	} else if score >= 60 && score <= 69 {
		return "D"
	} else if score >= 0 && score <= 59 {
		return "F"
	} else {
		return "Invalid score"
	}
}
