package calc

import "errors"
import "github.com/fatih/color"

func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		errorMessage := color.RedString("Provide more than 2 numbers and try again.")
		return errors.New(errorMessage), sum
	} else {
		for _, num := range numbers {
			sum += num
		}
		return nil, sum
	}
}
