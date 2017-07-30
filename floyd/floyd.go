package floyd

//Triangle Output flyd's triangle with given number of rows .
func Triangle(rows int) [][]int {

	output := make([][]int, rows)

	//the counter
	i := 1

	// iterate through rows
	for row := 0; row < rows; row++ {
		output[row] = make([]int, row + 1)

		// iterate columns
		for col := 0; col <= row; col++ {
			output[row][col] = i
			i++
		}
	}

	return output
}
