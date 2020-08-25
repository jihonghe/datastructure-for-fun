package sort_algrithm

func ShellSort(nums []int) {
	for step := len(nums); step >= 1; step /= 2 {
		for i := step; i < len(nums); i += step {
			for j := i - step; j >= 0; j -= step {
				if nums[j] > nums[j + step] {
					nums[j], nums[j + step] = nums[j + step], nums[j]
				}
			}
		}
	}
}
