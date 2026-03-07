package hot100practice

import "sort"

func Merge(intervals [][]int) [][]int {
	// 二维数组 只需要比较row行即可
	// 先对数组的每一行的start排序，将数组行排成从小到大
	// 循环遍历 判断后一行 的start是否小于等于当前行的end 如果小于等于则说明有重复
	// 然后就标记上从数组哪个位置开始有重复，然后就接着判断下一行与上一行是否也存在重复，
	// 当不存在重复的时候 就将标记开始的数组索引的第一个值与结束的行的第二个值，组成一个数组放到最终结果
	// 开始数组索引置为-1
	// 如果后一行 大于当前行 就将当前行直接放到数组中
	// 直到循环结束
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))
	// curStart, curEnd 用来记录当前连续的数组的起始和结束
	curStart, curEnd := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		s, e := intervals[i][0], intervals[i][1]
		if s <= curEnd {
			// 这里判断是因为防止后一行的end小于前一行的end
			// 比如 [1, 4] [2, 3]
			// 所以 需要将e赋值给curEnd 以保证curEnd是比较大的那个
			if e > curEnd {
				curEnd = e
			}
		} else {
			res = append(res, []int{curStart, curEnd})
			curStart, curEnd = s, e
		}
	}
	res = append(res, []int{curStart, curEnd})
	return res
}
