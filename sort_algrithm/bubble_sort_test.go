package sort_algrithm

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
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

func TestBubbleSort(t *testing.T) {
	length := 10
	min, max := -100, 100
	nums := generateNums(length, min, max)
	if suc, expectedNums, sortedNums := checkSortFunc(nums, BubbleSort); !suc {
		t.Errorf("BubbleSort failed:\nSource nums: %v\nexpected sort result: %v\nbubble sort result:%v\n",
			nums, expectedNums, sortedNums)
	} else {
		t.Logf("Pass BubbleSort：\nsource-nums: %v\nexpected-result: %v\nbubble-sort-result:%v\n",
			nums, expectedNums, sortedNums)
	}
}

func TestBubbleSortOptimized(t *testing.T) {
	length := 10
	min, max := -100, 100
	nums := generateNums(length, min, max)
	if suc, expectedNums, sortedNums := checkSortFunc(nums, BubbleSortOptimized); !suc {
		t.Errorf("BubbleSort failed:\nSource nums: %v\nexpected sort result: %v\nbubble sort result:%v\n",
			nums, expectedNums, sortedNums)
	} else {
		t.Logf("Pass BubbleSort：\nsource-nums: %v\nexpected-result: %v\nbubble-sort-result:%v\n",
			nums, expectedNums, sortedNums)
	}
}

func TestDoubleBubbleSort(t *testing.T) {
	length := 10
	min, max := -100, 100
	nums := generateNums(length, min, max)
	if suc, expectedNums, sortedNums := checkSortFunc(nums, DoubleBubbleSort); !suc {
		t.Errorf("BubbleSort failed:\nSource nums: %v\nexpected sort result: %v\nbubble sort result:%v\n",
			nums, expectedNums, sortedNums)
	} else {
		t.Logf("Pass BubbleSort：\nsource-nums: %v\nexpected-result: %v\nbubble-sort-result:%v\n",
			nums, expectedNums, sortedNums)
	}
}
