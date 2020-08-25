package sort_algrithm

func InsertionSort(nums []int) {
	// i指向待查找排序的元素的下标，默认从数组的第二个元素开始
	for i := 1; i < len(nums); i++ {
		// 将该元素移动到有序子数组的正确位置
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[j + 1] {
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
			}
		}
	}
}
