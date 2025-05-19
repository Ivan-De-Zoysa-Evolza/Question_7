package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	for {
		principal := getValidFloatInput("\nEnter the loan principal amount: ")
		annualRate := getValidFloatInput("Enter the annual interest rate (in %): ")
		term := getValidFloatInput("Enter the loan term (in months): ")

		monthlyPayment := calculateMonthlyPayment(principal, annualRate, int(term))
		fmt.Printf("\nYour monthly payment is: %.2f\n", monthlyPayment)

		fmt.Print("\nDo you want to calculate another loan? (yes/no): ")
		var again string
		fmt.Scanln(&again)
		if strings.ToLower(again) != "yes" {
			fmt.Println("\nGoodbye!")
			break
		}
	}
}

// Input validation: must be a number with up to 2 decimal places
func getValidFloatInput(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("\nInvalid input. Please enter a valid number.")
			continue
		}

		// Check decimal places
		if strings.Contains(input, ".") {
			parts := strings.Split(input, ".")
			if len(parts[1]) > 2 {
				fmt.Println("\nPlease enter a number")
				continue
			}
		}

		return value
	}
}

// Formula to calculate monthly payment
func calculateMonthlyPayment(P float64, annualRate float64, months int) float64 {
	if months == 0 {
		fmt.Println("\nError: Loan term must be greater than 0 months.")
		return 0
	}

	r := annualRate / 12 / 100
	if r == 0 {
		return P / float64(months) // No interest
	}

	numerator := P * r * math.Pow(1+r, float64(months))
	denominator := math.Pow(1+r, float64(months)) - 1
	return numerator / denominator
}
