package part_one

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str) // str是字符串类型 不可修改 需要改成byte数组类型
		// 通过该函数就能将byte数组中的字符按照从小到大来排序
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		// 将s改为string类型
		sortedStr := string(s)
		// 将str加入到mp[sortedStr]这个字符串数组中 然后再赋值给mp[sortedStr]
		// mp[sortedStr]也就是字符串排序后为sortedStr对应的字符串数组
		// 这样 结果就是key为sortedStr value为所有排序后和sortedStr相同的字符串数组
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
