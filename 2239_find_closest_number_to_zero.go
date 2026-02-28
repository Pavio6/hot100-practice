package hot100practice

// 找到离 0 最近的数字
func FindClosestNumber(nums []int) int {
    // 遍历每一个数 如果当前数的绝对值 小于 closest，则将closest的值赋值为当前值，i为当前值的索引，
    // 如果当前数的绝对值与closest相等，则取出该值的索引，
    // 和closest的索引，在原数组中看哪个值大，就赋值给closest
    cValue := 100000000
    cIndex := 0
    for i, num := range nums {
        if abs(num) < cValue {
            cValue = abs(num)
            cIndex = i
        } else if (abs(num) == cValue) {
            if (nums[i] > nums[cIndex]) {
                cIndex = i
            }
        }
        
    }
    return nums[cIndex]
}

// 最优解法
func ndClosestNumber02(nums []int) int {
    best := nums[0]
	// 不需要保存index，只需要比较当前数和best哪个更接近0，如果相等则比较哪个更大
    for _, x := range nums[1:] {
        if abs(x) < abs(best) || (abs(x) == abs(best) && x > best) {
            best = x
        }
    }
    return best
}


func abs(x int) int {
    if (x < 0) {
        return -x
    }
    return x
}