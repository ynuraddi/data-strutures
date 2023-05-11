package main

func main() {
	isValidSudoku(nil)
}

func isValidSudoku(board [][]byte) bool {
	ssx := make([]map[byte]struct{}, 9)
	ssy := make([]map[byte]struct{}, 9)
	ssb := make([]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		ssx[i] = make(map[byte]struct{}, 9)
		ssy[i] = make(map[byte]struct{}, 9)
		ssb[i] = make(map[byte]struct{}, 9)
	}

	count := 0

	for x, xv := range board {
		for y, v := range xv {
			switch {
			case v != byte('.') && x < 3 && y < 3:
				if _, ok := ssb[0][v]; ok {
					return false
				} else {
					ssb[0][v] = struct{}{}
				}
			case v != byte('.') && x < 3 && y < 6:
				if _, ok := ssb[1][v]; ok {
					return false
				} else {
					ssb[1][v] = struct{}{}
				}
			case v != byte('.') && x < 3 && y < 9:
				if _, ok := ssb[2][v]; ok {
					return false
				} else {
					ssb[2][v] = struct{}{}
				}
			case v != byte('.') && x < 6 && y < 3:
				if _, ok := ssb[3][v]; ok {
					return false
				} else {
					ssb[3][v] = struct{}{}
				}
			case v != byte('.') && x < 6 && y < 6:
				if _, ok := ssb[4][v]; ok {
					return false
				} else {
					ssb[4][v] = struct{}{}
				}
			case v != byte('.') && x < 6 && y < 9:
				if _, ok := ssb[5][v]; ok {
					return false
				} else {
					ssb[5][v] = struct{}{}
				}
			case v != byte('.') && x < 9 && y < 3:
				if _, ok := ssb[6][v]; ok {
					return false
				} else {
					ssb[6][v] = struct{}{}
				}
			case v != byte('.') && x < 9 && y < 6:
				if _, ok := ssb[7][v]; ok {
					return false
				} else {
					ssb[7][v] = struct{}{}
				}
			case v != byte('.') && x < 9 && y < 9:
				if _, ok := ssb[8][v]; ok {
					return false
				} else {
					ssb[8][v] = struct{}{}
				}
			}

			if _, ok := ssx[x][v]; ok {
				return false
			}

			if _, ok := ssy[y][v]; ok {
				return false
			}

			if v != byte('.') {
				ssx[x][v] = struct{}{}
				ssy[y][v] = struct{}{}
				count++
			}
		}
	}

	return true
}
