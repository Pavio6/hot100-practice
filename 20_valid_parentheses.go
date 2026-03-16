package hot100practice

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
// 示例 1：
//
// 输入：s = "()"
// 输出：true
// 示例 2：
//
// 输入：s = "()[]{}"
// 输出：true
// 示例 3：
//
// 输入：s = "(]"
// 输出：false
func IsValid(s string) bool {
    stack := []byte{}
    m := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{', 
    }
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c == '(' || c == '[' || c == '{' {
            stack = append(stack, c)
        } else {
            // 说明stack起始位置就是闭括号，直接返回false
            if len(stack) == 0 {
                return false
            }
            // 获取当前stack的顶部
            top := stack[len(stack)-1]
            // 说明顶部和当前字符抵消不了
            if top != m[c] {
                return false
            }
            // 抵消之后 将stack截取
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}