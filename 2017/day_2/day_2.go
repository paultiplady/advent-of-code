package day_2

import (
	"strings"
	"strconv"
	"fmt"
)

func calculateRow(row string) int {
	digits := strings.Fields(row)

	max := 0
	// If we get input larger than this, the calculation will break.
	// TODO: Use infinity here? E.g. https://blog.golang.org/constants
	min := 100000

	for _, char := range digits {
		digit, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Processing digit: %d\n", digit)
		if digit < min {
			min = digit
			fmt.Printf("New min: %d\n", min)
		}
		if digit > max {
			max = digit
			fmt.Printf("New max: %d\n", max)
		}
	}
	difference := max - min
	fmt.Printf("Difference: %d\n", difference)
	return difference
}

func calculateFile(file string) int {
	rows := strings.Split(file, "\n")

	total := 0

	for _, row := range rows {
		total += calculateRow(row)
	}
	return total
}