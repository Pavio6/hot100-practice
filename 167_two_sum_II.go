package hot100practice

// 双指针
func TwoSumII(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1} // 返回的是索引+1
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}