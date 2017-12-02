package day_1

import (
	"fmt"
)

func calculate(input string) int {
	total := 0

	for i, char := range input {
		digit := int(char - '0')
		fmt.Printf("Processing digit: %d\n", digit)

		j := (i + 1) % len(input)
		fmt.Printf("Next index: %d\n", j)
		nextDigit := int(input[j] - '0')

		fmt.Printf("Next digit: %d\n", nextDigit)

		if digit == nextDigit {
			fmt.Printf("Repeated digit, adding to total: %d\n", digit)
			total += digit
		}
	}
	return total
}

func main() {

}