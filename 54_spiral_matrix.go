package hot100practice

// 54. 螺旋矩阵 
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
// 示例 1：
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
// 示例 2：

// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]

// 提示：

// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
func SpiralOrder(matrix [][]int) []int {
	// 从左到右（顶行）
	// 从上到下（右列）
	// 从右到左（底行）
	// 从下到上（左列）
	// 一次循环就是上述四个步骤，重复循环即可
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	row, column := len(matrix), len(matrix[0])
	top, bottom := 0, row-1
	left, right := 0, column-1
	res := make([]int, 0, row*column)

	for left <= right && top <= bottom {
		// top row
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++
		// right col
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top > bottom || left > right {
			break
		}
		// bottom row
		for j := right; j >= left; j-- {
			res = append(res, matrix[bottom][j])
		}
		bottom--
		// left col
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res

}
