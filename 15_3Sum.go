package hot100practice

import "sort"


func ThreeSum(nums []int) [][]int {

	// 将数组从小到大排序
	sort.Ints(nums)

	res := [][]int{}
	// i < len(nums)-2 原因是i后面必须有两个数
	// [-1, 0, 1, 2, -1, -4]
	// [-4, -1, -1, 0, 1, 2]
	for i := 0; i < len(nums)-2; i++ {

		// 去重
		// i > 0 为了跳过第一次比较，当i=0时
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针进行遍历
		l := i + 1
		r := len(nums) - 1

		for l < r {

			sum := nums[i] + nums[l] + nums[r]

			if sum == 0 {

				res = append(res, []int{nums[i], nums[l], nums[r]})

				// 去重
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--

			} else if sum < 0 {
				l++
			} else {
				r--
			}

		}

	}

	return res
}
