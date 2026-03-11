package hot100practice

// 977. 有序数组的平方
// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

// 示例 1：

// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]
// 排序后，数组变为 [0,1,9,16,100]
// 示例 2：

// 输入：nums = [-7,-3,2,3,11]
// 输出：[4,9,9,49,121]
// 解释：平方后，数组变为 [49,9,4,9,121]
// 排序后，数组变为 [4,9,9,49,121]
func SortedSquares(nums []int) []int {
    // 最大的平方一定来源于数组的两端的值
    // 从后往前遍历 每次都取最大值填充到位置i
    n := len(nums)
    res := make([]int, n)
    l := 0
    r := n-1
    for i := n - 1; i >= 0; i-- {
        left := nums[l] * nums[l]
        right := nums[r] * nums[r]
        if left > right {
            res[i] = left
            l++
        } else {
            res[i] = right
            r--
        }
    }
    return res
}