package chessboard

// File type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Chessboard type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

func countOccupiedInFile(file File) int {
	counter := 0

	for _, occupied := range file {
		if occupied {
			counter++
		}
	}

	return counter
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {

	fileRow, existsFile := cb[file]

	if !existsFile {
		return 0
	}

	return countOccupiedInFile(fileRow)
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {

	if rank <= 0 {
		return 0
	}

	counter := 0
	for _, file := range cb {

		if rank > len(file) {
			continue
		}

		if file[rank-1] {
			counter++
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	counter := 0
	for _, file := range cb {
		counter += len(file)
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	counter := 0
	for _, file := range cb {

		counter += countOccupiedInFile(file)
	}
	return counter
}
