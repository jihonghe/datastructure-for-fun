package sort_algrithm

import (
	"fmt"
	"math/rand"
	"sort"
)

func generateNums(length, min, max int) []int {
	if min >= max {
		panic(fmt.Sprintf("Error: min(%d) is greater than max(%d)\n", min, max))
	}

	nums := make([]int, length)
	for i := 0; i < length; i++ {
		nums[i] = rand.Intn(max-min) + min
	}

	return nums
}

func checkSortFunc(nums []int, sortFunc func([]int)) (suc bool, expectedNums, sortedNums []int) {
	expectedNums = make([]int, len(nums))

	copy(expectedNums, nums)

	sort.Ints(expectedNums)
	sortFunc(nums)
	sortedNums = nums

	for i, num := range nums {
		if expectedNums[i] != num {
			return false, expectedNums, sortedNums
		}
	}

	return true, expectedNums, sortedNums
}
