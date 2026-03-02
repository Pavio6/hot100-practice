package hot100practice

import "strconv"

// 题目：总结区间（LeetCode 228）
// 给定一个无重复元素的有序整数数组 nums 。返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。
// 也就是说，nums 中的每个元素都被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。列表中的每个区间范围 [a,b] 应该按如下格式输出：
// "a->b" ，如果 a != b
// "a" ，如果 a == b
func SummaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	var res []string
	var head int = 0

	for i := range nums {
		if i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			continue
		}
		if head == i {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			tmp := strconv.Itoa(nums[head]) + "->" + strconv.Itoa(nums[i])
			res = append(res, tmp)
		}

		head = i + 1
	}

	return res
}
