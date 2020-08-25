package sort_algrithm

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	length := 10
	min, max := -100, 100
	nums := generateNums(length, min, max)
	testFunc := BubbleSort

	if suc, expectedNums, sortedNums := checkSortFunc(nums, testFunc); !suc {
		t.Logf("\nsource-nums: %v\nexpected-result: %v\nactual-result:%v\n",
			nums, expectedNums, sortedNums)
	} else {
		t.Errorf("\nsource-nums: %v\nexpected-result: %v\nactual-result:%v\n",
			nums, expectedNums, sortedNums)
	}
}

func TestBubbleSortOptimized(t *testing.T) {
	length := 10
	min, max := -100, 100
	nums := generateNums(length, min, max)
	testFunc := BubbleSortOptimized

	if suc, expectedNums, sortedNums := checkSortFunc(nums, testFunc); !suc {
		t.Logf("\nsource-nums: %v\nexpected-result: %v\nactual-result:%v\n",
			nums, expectedNums, sortedNums)
	} else {
		t.Errorf("\nsource-nums: %v\nexpected-result: %v\nactual-result:%v\n",
			nums, expectedNums, sortedNums)
	}
}

func TestDoubleBubbleSort(t *testing.T) {
	length := 10
	min, max := -100, 100
	nums := generateNums(length, min, max)
	testFunc := DoubleBubbleSort

	if suc, expectedNums, sortedNums := checkSortFunc(nums, testFunc); !suc {
		t.Errorf("\nsource-nums: %v\nexpected-result: %v\nactual-result:%v\n",
			nums, expectedNums, sortedNums)
	} else {
		t.Logf("\nsource-nums: %v\nexpected-result: %v\nactual-result:%v\n",
			nums, expectedNums, sortedNums)
	}
}
