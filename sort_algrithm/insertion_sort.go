package sort_algrithm

func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j - 1] {
				nums[j], nums[j - 1] = nums[j -1], nums[j]
			}
		}
	}
}
