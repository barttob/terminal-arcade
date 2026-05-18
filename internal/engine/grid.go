package engine

type Point struct {
	X int
	Y int
}

type Grid struct {
	Width  int
	Height int
	Cells  [][]rune
}
