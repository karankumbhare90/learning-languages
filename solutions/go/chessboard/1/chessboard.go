package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
// File represents a column of squares (true if occupied)
type File []bool

// Chessboard maps file names ("A"-"H") to their respective columns
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, ok := cb[file];

    if !ok {
        return 0;
    }

    count := 0;
    for _, occupied := range f {
        if occupied {
            count++
        }
    }

    return count;
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
        return 0
    }

    count := 0
    index := rank - 1 // rank 1 -> index 0, rank 8 -> index 7

    for _, file := range cb {
        if index < len(file) && file[index] {
            count++
        }
    }

    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
    for _, file := range cb {
        count += len(file)
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
    for _, file := range cb {
        for _, occupied := range file {
            if occupied {
                count++
            }
        }
    }
    return count
}
