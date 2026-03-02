package hot100practice

import "sort"

// 题目：最长公共前缀（LeetCode 14）
// 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
func LongestCommonPrefix(strs []string) string {
	// 1.strs长度为0，直接返回空字符串
	// 2.以strs[0]为prefix，循环判断后续每一个str是否包含prefix
	// 针对每一个str，循环对其进行判断，如果不包含，则将prefix从最后截断一位，再次判断，
	// 直到str包含prefix，前提prefix的长度要大于0，如果在这过程中=0了，则直接返回空
	// 3.当对一个str判断之后，在外围for循环接着判断下一个str，同样的方法
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		s := strs[i]
		for len(prefix) > 0 && !hasPrefix(s, prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}
func hasPrefix(s, p string) bool {
	if len(p) > len(s) {
		return false
	}
	for i := 0; i < len(p); i++ {
		if s[i] != p[i] {
			return false
		}
	}
	return true
}

// 排序法
func LongestCommonPrefix02(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	// 会将strs中的字符串从小到大按字典序排序，所以最左边就是最小的字符串，最右边就是最大的
	sort.Strings(strs)

	first := strs[0]
	last := strs[len(strs)-1]
	// 判断从开始到最后有几个字符相等 最后通过[:i]获得相同的prefix
	i := 0
	for i < len(first) && i < len(last) && first[i] == last[i] {
		i++
	}

	return first[:i]
}
