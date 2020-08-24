package sort_algrithm

/**
待排序数组：[10, -1, 5, 7, 1, 0, 4, 8]
思路：双循环处理，冒泡方向为索引递增方向。外循环用于界定已冒泡与未冒泡的边界，内循环用于大小比较及位置交换。
*/
func BubbleSort(nums []int) {
	// i为当前冒泡的目标位置
	for i := len(nums) - 1; i > 0; i-- {
		// j为冒泡的起始位置
		for j := 0; j < i; j++ {
			if nums[j] >= nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

/**
改进：冒泡排序中有两步，一是比较大小，二是交换位置（冒泡）。在经典的冒泡排序中，即便无需交换位置，仍会走完O(n2)的比较过程。而对于已经有序的数组，
我们实际是可以不用再话O(n2)的时间复杂度去遍历比较了。因此，我们可以加入一个标志位。判断在一轮的冒泡中，如果没有发生位置交换，那就说明数组有序，
可以直接结束排序。
*/
func BubbleSortOptimized(nums []int) {
	var exchanged bool

	for i := len(nums) - 1; i > 0; i-- {
		exchanged = false
		// j为冒泡的起始位置
		for j := 0; j < i; j++ {
			if nums[j] >= nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				exchanged = true
			}
		}

		if !exchanged {
			break
		}
	}
}

func DoubleBubbleSort(nums []int) {
	var exchanged bool
	endIndex := len(nums) - 1

	for i := endIndex; i > 0; i-- {
		exchanged = false
		// 大的往右，小的往左
		for l, r := endIndex-i, i; l < i; l, r = l+1, r-1 {
			if nums[l] > nums[l+1] {
				nums[l], nums[l+1] = nums[l+1], nums[l]
				exchanged = true
			}

			if nums[r] < nums[r-1] {
				nums[r], nums[r-1] = nums[r-1], nums[r]
				exchanged = true
			}
		}

		if !exchanged {
			break
		}
	}
}
