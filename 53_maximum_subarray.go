package hot100practice

// 最大连续子数组和
// 使用Kadane算法（动态规划），时间复杂度O(n)，空间复杂度O(1)
func MaxSubArray(nums []int) int  {
	max, curSum := nums[0], 0
	for _, num := range nums {
		curSum += num
		if curSum > max {
			max = curSum
		}
		if curSum < 0 {
			curSum = 0
		}
	}
	return max
}