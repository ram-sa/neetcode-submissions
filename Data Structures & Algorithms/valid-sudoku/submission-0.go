func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]bool)
	cols := make(map[int]map[byte]bool)
	qdts := make(map[int]map[byte]bool)

	for r, arr := range board {
		for c, val := range arr {
			//Ignore empty squares
			if val == '.' { continue }

			//Calculate quadrant
			q := (r/3) * 3 + (c/3)

			//fmt.Println(r," - ",c," - ",q," - ", string(val))
			
			//Initalize the row if needed
			if _, ok := rows[r]; !ok {
				rows[r] = make(map[byte]bool)
			}
			//Check if the number already exists in this row, add it otherwise
			if rows[r][val]{
				return false
			}
			rows[r][val] = true

			//Same as above but for columns
			if _, ok := cols[c]; !ok {
				cols[c] = make(map[byte]bool)
			}
			if cols[c][val]{
				return false
			}
			cols[c][val] = true

			//Same as above for quadrants
			if _, ok := qdts[q]; !ok {
				qdts[q] = make(map[byte]bool)
			}
			if qdts[q][val]{
				return false
			}
			qdts[q][val] = true
		}
	}

	return true
}
