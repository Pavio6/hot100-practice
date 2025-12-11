package part_one

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ { // 使用 len() 获取切片长度
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j} // 返回切片，而不是数组
			}
		}
	}
	return nil // 如果没有找到结果，返回 nil
}
