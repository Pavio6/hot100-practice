package hot100practice

// 题目：接雨水（LeetCode 42）
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

func Trap(height []int) int {
    // 方法：前缀最大值和后缀最大值 DP的思想

    n := len(height)
    if (n == 0) {
        return 0
    }
    lMax := make([]int, n)
    rMax := make([]int, n)
    res := 0
    // 思路: 1.先求出height[i]在i位置的左侧的最大值 比如i=3 那么就是求height[0],height[1],height[2]的最大值
    // 2.求出height[i]在位置i的右侧的最大值
    lMax[0] = height[0]
    for i := 1; i < n; i++ {
        if (height[i] > lMax[i-1]) {
            lMax[i] = height[i]
        } else {
            // 每一轮循环lMax[i]的值都是i左侧的最大值，所以如果当前i小于lMax[i-1]，
            // 就将当前lMax[i]的值赋值为lMax[i-1]
            lMax[i] = lMax[i-1]
        }
    }
    rMax[n-1] = height[n-1]
    // 这里 i 的判断条件一定要有 = 0 , 否则无法获取rMax[0]的值
    for i := n-2; i >= 0; i-- {
        if (height[i] > rMax[i+1]) {
            rMax[i] = height[i]
        } else {
            rMax[i] = rMax[i+1]
        }
    }
    for i := 0; i < n; i++ {
        water := min(lMax[i], rMax[i]) - height[i]
        res += water
    }
    return res
}
func min(a, b int) int {
    if (a > b) {
        return b
    }
    return a
}