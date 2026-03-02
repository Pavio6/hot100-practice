package hot100practice

// 题目：两数之和 II - 输入有序数组（LeetCode 167）
// 给定一个按非递减顺序排列的整数数组 numbers 和目标值 target，
// 请在数组中找出两个数，使它们的和等于 target，并返回它们的下标（从 1 开始计数）。
// 题目保证恰好存在一个有效答案，且同一个元素不能重复使用。
// 常见解法是双指针：左指针从头开始，右指针从尾开始，根据当前和与 target 的大小关系移动指针。
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
