package hot100practice

// 题目：交替合并字符串（LeetCode 1768）
// 给你两个字符串 word1 和 word2，请你按字符交替合并字符串。如果一个字符串比另一个短，
// 就在交替过程中尽可能多地使用较长字符串的字符。返回合并后的字符串。
func MergeAlternately(word1 string, word2 string) string {
    // 使用双指针解决
    // 初始化左右指针，左指针对应word1，右指针对应word2
    // for循环条件为左指针小于word1的长度或者右指针小于word2的长度
    // 向s中append字符，对应指针加一
    // 返回字符串s
    l, r := 0, 0
    len1, len2 := len(word1), len(word2)
    s := make([]byte, 0, len1+len2)
    for l < len1 || r < len2 {
        if l < len1 {
            s = append(s, word1[l])
            l++
        }
        if r < len2 {
            s = append(s, word2[r])
            r++
        }
    }
    return string(s)
}