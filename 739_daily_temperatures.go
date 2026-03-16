package hot100practice

func DailyTemperatures(temperatures []int) []int {

    n := len(temperatures)

    ans := make([]int, n)
    // 栈中存还没有找到更高温度的 index
    stack := []int{}

    for i := 0; i < n; i++ {
        // 当前温度大于栈顶温度 将栈顶温度弹出
        // 也就是帮栈顶的那一天找到了温度
        // for循环意思就是不断的帮很多天找到答案
        for len(stack) > 0 &&
            temperatures[i] > temperatures[stack[len(stack)-1]] {
            // 拿到栈顶元素
            prev := stack[len(stack)-1]
            // pop
            stack = stack[:len(stack)-1]
            // 计算等待天数
            ans[prev] = i - prev
        }

        stack = append(stack, i)
    }

    return ans
}