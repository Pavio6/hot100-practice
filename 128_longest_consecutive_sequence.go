package hot100practice

// 128. 最长连续序列
// 给定一个未排序的整数数组，找出最长连续序列的长度。
// 要求算法的时间复杂度为 O(n)。
//
// 示例:
//
// 输入: [100, 4, 200, 1, 3, 2]
// 输出: 4
// 解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
func LongestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for _, num := range nums {
        set[num] = struct{}{}
    }
    longest := 0

    for num := range set {
        // num-1 不存在 说明是序列起点
        if _, ok := set[num-1]; !ok {
            cur := num
            l := 1
            for {
                if _, ok := set[cur+1]; ok {
                    cur++
                    l++
                } else {
                    break
                }
            }
            if l > longest {
                longest = l
            }
        }
    }
    return longest
}