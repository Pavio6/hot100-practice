package hot100practice

// 题目：除自身以外数组的乘积（LeetCode 238）
// 给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
// 说明：请不要使用除法，且在 O(n) 时间复杂度内完成此题。

// 前缀积 + 后缀积
func ProductExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)

    // 1.先完成索引为i的左侧的乘积，比如res[2]，这一步就会得到res[0]*res[1]
    // 后续会再得到res[2]右侧的乘积
    prefix := 1
    for i := 0; i < n; i++ {
        res[i] = prefix
        prefix *= nums[i]
    }
    // 2.完成索引为i的右侧的乘积，比如res[2]，这一步会得到res[3]...这些值，然后再乘以之前的res的值(左侧乘积)
    // 就完成了题目要求
    suffix := 1
    for i := n - 1; i >= 0; i-- {
        res[i] *= suffix
        suffix *= nums[i]
    }

    return res
}