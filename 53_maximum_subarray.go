package hot100practice

// 最大连续子数组和
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