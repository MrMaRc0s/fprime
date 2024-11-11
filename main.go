package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the number of arguments is exactly 1 (program name + 1 argument)
	if len(os.Args) != 2 {
		return
	}

	// Parse the argument as an integer
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 1 {
		return
	}

	// Find and print the prime factors
	var result []int
	for i := 2; i*i <= num; i++ {
		for num%i == 0 {
			result = append(result, i)
			num /= i
		}
	}
	if num > 1 {
		result = append(result, num)
	}

	// Print the result in the required format
	for i, factor := range result {
		if i > 0 {
			fmt.Print("*")
		}
		fmt.Print(factor)
	}
	fmt.Println()
}
