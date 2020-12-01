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

	// The above solution is certainly not the fastest mechanism. We _should_ walk inwards from
	// both sides of a sorted array. As these numbers are already sorted, we
	// should be able to take a look at some specified indexes, starting with 0
	// and `len(numbers)-1`. If the sum is too large, then we walk decrement the
	// upper index. If the sum is too small, we increment the lower index

	// Example:
	// target: 59
	// numbers: [ 1, 2, 5, 19, 34, 46, 52, 67]
	// i = 0; j = 1; k = 7; numbers[i] = 1; numbers[j] = 2; numbers[k] = 67; sum(i,j,k) = 70 => decrement k
	// i = 0; j = 1; k = 6; numbers[i] = 1; numbers[j] = 2; numbers[k] = 52; sum(i,j,k) = 55 => sum(j,k) = 54 => increment j
	// i = 0; j = 2; k = 6; numbers[i] = 1; numbers[j] = 5; numbers[k] = 52; sum(i,j,k) = 58 => sum(j,k) = 53 => increment i
	// i = 1; j = 2; k = 6; numbers[i] = 2; numbers[j] = 5; numbers[k] = 52; sum(i,j,k) = 59 => MATCH

	sort.Ints(numbers)

	i := 0
	j := 1
	k := len(numbers) - 1
	for {
		if i >= j || j >= k {
			// This should not happen with a good data set.
			return -1, fmt.Errorf("Invalid data set")
		}

		x0 := numbers[i]
		x1 := numbers[j]
		x2 := numbers[k]
		if x0+x1+x2 == target {
			return x0 * x1 * x2, nil
		} else if x0+x1+x2 > target {
			k--
		} else if x1+x2 < target {
			j++
		} else {
			i++
		}
	}
}
