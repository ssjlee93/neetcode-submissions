

type NumMatrix struct {
	P [][]int // prefix sum. P for prefix
}

func Constructor(matrix [][]int) NumMatrix {
	prefixSum := make([][]int, 0, len(matrix))
	init := NumMatrix{prefixSum}
	for _, row := range matrix {
		prefixRow := make([]int, 0, len(matrix[0]))
		for j, n := range row {
			if j == 0 {
				prefixRow = append(prefixRow, n)
				continue 
			}
			prefixRow = append(prefixRow, prefixRow[len(prefixRow)-1] + n)
		}
		init.P = append(init.P, prefixRow)
	}
	return init
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	rowFrom := min(row1, row2)
	rowTo := max(row1, row2)
	colFrom := min(col1, col2)
	colTo := max(col1, col2)
	ans := 0
	for i := rowFrom; i <= rowTo; i++ {
		prefix := this.P[i][colTo]
		if colFrom > 0 {
			prefix -= this.P[i][colFrom-1]
		}
		ans += prefix
	}
	return ans
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)

// edge : can sumRegion gimme upper right and lower left corner ?
// if so, we need to do min max on row, col to determine how far we go.

// should i do prefix sum for each row or for all ? 
// i think each row should be easier to pull out numbers
// then when we have a region, we loop for each row from row1 to row2
// for each iter, we pull out prefixSum of max col1, col2 and sub min col1, col2 - 1 