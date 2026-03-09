package hot100practice

// 217. 存在重复元素
// 给你一个整数数组 nums ，判断是否存在重复元素。如果存在一值在数组中出现至少两次，返回 true ；如果数组中每个元素都不相同，返回 false 。
// 示例 1：

// 输入：nums = [1,2,3,1]
// 输出：true
// 示例 2：

// 输入：nums = [1,2,3,4]
// 输出：false
// 示例 3：

// 输入：nums = [1,1,1,3,3,4,3,2,4,2]
// 输出：true
func ContainsDuplicate(nums []int) bool {
	// 为map提前分配容量，避免扩容带来的性能损失
	// value使用空结构体 struct{}，因为它不占用内存
    seen := make(map[int]struct{}, len(nums))

    for _, v := range nums {
        if _, ok := seen[v]; ok {
            return true
        }
        seen[v] = struct{}{}
    }

    return false
}