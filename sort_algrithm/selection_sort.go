package sort_algrithm

func SelectionSort(nums []int) {
	endIndex := len(nums) - 1
	for i := 0; i < endIndex; i++ {
		minIndex := i
		for j := i + 1; j <= endIndex; j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}

		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

// 改进：常规的选择排序是单向进行的，我们可以双向选择。例如：按升序排序，我们可以在内循环中同时选择出局部最小值和局部最大值，然后分别插入两边
func DoubleSelectionSort(nums []int) {
	endIndex := len(nums) - 1
	for left, right := 0, endIndex; left < right; left, right = left + 1, right - 1 {
		minIndex := left
		maxIndex := right
		for j := left; j <= right; j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
			if nums[maxIndex] < nums[j] {
				maxIndex = j
			}
		}

		nums[left], nums[minIndex] = nums[minIndex], nums[left]
		nums[right], nums[maxIndex] = nums[maxIndex], nums[right]
	}
}
