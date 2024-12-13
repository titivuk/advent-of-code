package utils

var (
	StraightDirections = [4][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
)

type Point struct {
	X, Y int
}

func InRange(m, n, i, j int) bool {
	return 0 <= i && i < m && 0 <= j && j < n
}

func IsDigit(ch byte) bool {
	return 48 <= ch && ch <= 57
}
