package hot100practice

// 125. 验证回文串
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。

func IsPalindrome(s string) bool {
   
    cArr := []rune{}
    // 遍历字符串s时，每一个字符都是rune类型，所以需要使用rune类型数组来接收
    for _, c := range s {
        if c >= 'A' && c <= 'Z' {
            cArr = append(cArr, c+32)
        } 
        if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
            cArr = append(cArr, c)
        } 
    }
    
    sNew := string(cArr)
  
    l := 0
    r := len(sNew)-1
   
    for l <= r {
        if (sNew[l] != sNew[r]) {
            return false
        }
        l++
        r--
    }
    return true
}