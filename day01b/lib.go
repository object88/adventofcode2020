package day01b

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

// Process looks for three numbers from the input that sum to the target
func Process(in io.Reader, target int) (int, error) {
	var numbers []int

	// We are going to parse & tokenize all in one go.
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return -1, err
		}
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return processNumbers(numbers, target)
}

func processNumbers(numbers []int, target int) (int, error) {
	// This technique operates on very rough O(n^3). This is written for
	// simplicity.
	// for i := 0; i < len(numbers); i++ {
	// 	for j := len(numbers) - 1; j >= 0; j-- {
	// 		for k := i; k < j; k++ {
	// 			if numbers[i]+numbers[j]+numbers[k] == target {
	// 				// We have found the desired pair; return their multiple.
	// 				return numbers[i] * numbers[j] * numbers[k], nil
	// 			}
	// 		}
	// 	}
	// }

	// return -1, fmt.Errorf("Failed to find combination")

	// The above solution is certainly not the fastest mechanism. We can loop
	// over a part of the array from both sides, while holding a third index
	// steady.

	sort.Ints(numbers)

	for i := 0; i < len(numbers); i-- {
		x0 := numbers[i]
		target0 := target - x0

		j := i + 1
		k := len(numbers) - 1
		for {
			if i == j {
				break
			}
			x1 := numbers[j]
			x2 := numbers[k]
			if x1+x2 == target0 {
				return x0 * x1 * x2, nil
			} else if x1+x2 > target0 {
				k--
			} else {
				j++
			}
		}
	}

	return -1, fmt.Errorf("failed")
}
