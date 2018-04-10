/* https://leetcode.com/problems/range-sum-query-2d-immutable/description/
Given a 2D matrix matrix, find the sum of the elements inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
https://leetcode.com/static/images/courses/range_sum_query_2d.png
The above rectangle (with the red border) is defined by (row1, col1) = (2, 1) and (row2, col2) = (4, 3), which contains sum = 8.

Example:

    Given matrix = [
      [3, 0, 1, 4, 2],
      [5, 6, 3, 2, 1],
      [1, 2, 0, 1, 5],
      [4, 1, 0, 1, 7],
      [1, 0, 3, 0, 5]
    ]

    sumRegion(2, 1, 4, 3) -> 8
    sumRegion(1, 1, 2, 2) -> 11
    sumRegion(1, 2, 2, 4) -> 12

Note:

    1. You may assume that the matrix does not change.
    2. There are many calls to sumRegion function.
    3. You may assume that row1 ≤ row2 and col1 ≤ col2.
*/

package leetcode

type NumMatrix struct {
	matrix [][]int
}

func NumMatrixConstructor(matrix [][]int) NumMatrix {
	row, col := len(matrix), len(matrix[0])
	tmp := make([][]int, row)

	// init idx==0 row
	tmp[0] = make([]int, col, col)
	tmp[0][0] = matrix[0][0]
	for j := 1; j < col; j++ {
		tmp[0][j] = tmp[0][j-1] + matrix[0][j]
	}

	// init 1 -- n rows
	for i := 1; i < row; i++ {
		tmp[i] = make([]int, col, col)
		tmp[i][0] = tmp[i-1][0] + matrix[i][0]
		for j := 1; j < col; j++ {
			tmp[i][j] = matrix[i][j] + tmp[i][j-1] + tmp[i-1][j] - tmp[i-1][j-1]
		}
	}

	return NumMatrix{matrix: tmp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.matrix[row2][col2] - this.matrix[row1-1][col2] - this.matrix[row2][col1-1] + this.matrix[row1-1][col1-1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
