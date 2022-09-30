package squares

type Point struct {
	x, y int
}

type Square struct {
	index      int
	coordinate Point
}

type Edge struct {
	squareIndex int
	edge        int
}
