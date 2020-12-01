package day01a

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

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

	// This technique operates on very rough O(n^2). This is written for
	// simplicity.
	// for i := 0; i < len(numbers); i++ {
	// 	for j := len(numbers) - 1; j >= 0; j-- {
	// 		if numbers[i]+numbers[j] == target {
	// 			// We have found the desired pair; return their multiple.
	// 			return numbers[i] * numbers[j], nil
	// 		}
	// 	}
	// }

	// This is certainly not the fastest mechanism. We _should_ walk inwards from
	// both sides of a sorted array. As these numbers are already sorted, we
	// should be able to take a look at some specified indexes, starting with 0
	// and `len(numbers)-1`. If the sum is too large, then we walk decrement the
	// upper index. If the sum is too small, we increment the lower index

	// Example:
	// target: 21
	// numbers: [ 1, 2, 5, 19, 34, 46]
	// i = 0; j = 5; numbers[i] = 1; numbers[j] = 46; sum = 47 => decrement j
	// i = 0; j = 4; numbers[i] = 1; numbers[j] = 34; sum = 35 => decrement j
	// i = 0; j = 3; numbers[i] = 1; numbers[j] = 19; sum = 20 => increment i
	// i = 1; j = 3; numbers[i] = 2; numbers[j] = 19; sum = 21 => MATCH

	sort.Ints(numbers)

	i := 0
	j := len(numbers) - 1
	for {
		if i >= j {
			// This should not happen with a good data set.
			return -1, fmt.Errorf("Invalid data set")
		}

		x := numbers[i] + numbers[j]
		if x == target {
			return numbers[i] * numbers[j], nil
		}
		if x > target {
			j--
		} else {
			i++
		}
	}
}
