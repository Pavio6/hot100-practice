package hot100practice

// 36. 有效的数独
// 请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
// 注意：
// 一个有效的数独（部分已被填充）不一定是可解的。
// 只需要根据以上规则，验证已经填入的数字是否有效即可。
// 空白格用 '.' 表示。
func IsValidSudoku(board [][]byte) bool {
	// 思路：创建三个二维数组，行，列，box
	// 每一行存储当前行中的所有的数，当遇到重复的时候，直接返回false，反之将当前行的当前数的值设为true
	// 三个二维数组都是这样
	var rows [9][9]bool
	var cols [9][9]bool
	var boxes [9][9]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '1'

			boxIndex := (i/3)*3 + j/3
			// cols 这里二维数组需要用列
			if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
				return false
			}
			rows[i][num] = true
			// #这里也是
			cols[j][num] = true
			boxes[boxIndex][num] = true
		}
	}
	return true
}
