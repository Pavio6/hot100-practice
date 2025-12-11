package part_one

func twoSum01(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ { // 使用 len() 获取切片长度
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j} // 返回切片，而不是数组
			}
		}
	}
	return nil // 如果没有找到结果，返回 nil
}

// 暴力枚举
// for .. range .. 用法 通常于遍历数组、切片、映射（map）、字符串
// i代表数组的索引 x代表该索引下的值
func twoSum02(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 最优解法 哈希表
func twoSum03(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		// hashTable[target-x] 如果存在该值 则返回该值在hash表中对应的key 赋值给p
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		// 没有找到 就先将这组数和下标放进哈希表中
		// key=数组下标 value=数组值
		hashTable[x] = i
	}
	return nil
}
