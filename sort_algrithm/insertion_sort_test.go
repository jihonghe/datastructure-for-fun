package sort_algrithm

import "testing"

func TestInsertionSort(t *testing.T) {
	length := 20
	min, max := -100, 100
	nums := generateNums(length, min, max)
	testFunc := InsertionSort

	if suc, expectedNums, sortedNums := checkSortFunc(nums, testFunc); !suc {
		t.Errorf("\nsource-nums: %v\nexpected-result: %v\nactual-result:%v\n",
			nums, expectedNums, sortedNums)
	} else {
		t.Logf("\nsource-nums: %v\nexpected-result: %v\nactual-result:%v\n",
			nums, expectedNums, sortedNums)
	}
}