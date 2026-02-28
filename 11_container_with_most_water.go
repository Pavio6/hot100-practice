package hot100practice

// 盛最多水的容器
func MaxArea(height []int) int {
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		h := height[l]
		if height[r] < h {
			h = height[r]
		}
		area := h * (r - l)
		if area > max {
			max = area
		}
		// 双指针移动较短的一边
		// 因为每次移动宽度都会减少，所以尽量保存较高的高度
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
