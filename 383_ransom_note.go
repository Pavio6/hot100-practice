package hot100practice

// 383. 赎金信
// 给定一个赎金信 (ransom) 字符串和一个杂志 (magazine) 字符串，判断第一个字符串 ransom 能不能由第二个字符串 magazines 里面的字符构成。如果可以构成，返回 true ；否则返回 false。
// (赎金信字符串和杂志字符串中的每个字符都是小写字母。)
// 示例 1：
// 输入：ransomNote = "a", magazine = "b"
// 输出：false
// 示例 2：
// 输入：ransomNote = "aa", magazine = "ab"
// 输出：false
// 示例 3：
// 输入：ransomNote = "aa", magazine = "aab"
// 输出：true
func CanConstruct(ransomNote string, magazine string) bool {

    if len(ransomNote) > len(magazine) {
        return false
    }

    r := map[rune]int{}
    m := map[rune]int{}

    for _, c := range ransomNote {
        r[c]++
    }

    for _, c := range magazine {
        m[c]++
    }

    for k, v := range r {
        if v > m[k] {
            return false
        }
    }

    return true
}