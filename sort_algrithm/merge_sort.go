package sort_algrithm

func MergeSort(nums []int) {
	mergesort(nums, 0, len(nums) - 1)
}

func mergesort(nums []int, left, right int) {
	if left == right {
		return
	}

	mid := left + (right - left) / 2
	mergesort(nums, left, mid)
	mergesort(nums, mid + 1, right)

	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	tmpNums := make([]int, 0, right - left + 1)

	leftIndex, rightIndex := left, mid + 1
	for leftIndex <= mid && rightIndex <= right {
		if nums[leftIndex] <= nums[rightIndex] {
			tmpNums = append(tmpNums, nums[leftIndex])
			leftIndex++
		} else {
			tmpNums = append(tmpNums, nums[rightIndex])
			rightIndex++
		}
	}
	for leftIndex <= mid {
		tmpNums = append(tmpNums, nums[leftIndex])
		leftIndex++
	}
	for rightIndex <= right {
		tmpNums = append(tmpNums, nums[rightIndex])
		rightIndex++
	}

	for i := 0; i < len(tmpNums); i++ {
		nums[left + i] = tmpNums[i]
	}
}
