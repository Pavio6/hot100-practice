package hot100practice

import "strconv"

func EvalRPN(tokens []string) int {
    // 依次遍历 
    // 如果是数字则入栈
    // 如果是符号 则将栈中前两个值弹出 计算出结果再入栈
    // 依次循环
    stack := []int{}
    for _, token := range tokens {
        if token == "+" || token == "-" || token == "*" || token == "/" {
            stack = pop(stack, token)
        } else {
            num, _ := strconv.Atoi(token)
            stack = push(stack, num)
        }
    }
    return stack[0]
}
func push(stack []int, v int) []int {
    return append(stack, v)
}

func pop(stack []int, oper string) []int {
    a := stack[len(stack)-2]
    b := stack[len(stack)-1]

    stack = stack[:len(stack)-2]
    var val int
    switch oper {
        case "+":
           val = a + b
        case "-":
            val = a - b
            
        case "*":
            val = a * b
            
        case "/":
            val = a / b
    }
    return append(stack, val)
}