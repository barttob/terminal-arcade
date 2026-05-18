package engine

type Direction Point

var (
	Up    = Direction{X: 0, Y: -1}
	Down  = Direction{X: 0, Y: 1}
	Left  = Direction{X: -1, Y: 0}
	Right = Direction{X: 1, Y: 0}
)

func DirectionFromKey(key string) (Direction, bool) {
	switch key {
	case "up", "w":
		return Up, true
	case "down", "s":
		return Down, true
	case "left", "a":
		return Left, true
	case "right", "d":
		return Right, true
	default:
		return Direction{}, false
	}
}

func (d Direction) Opposite() Direction {
	return Direction{X: -d.X, Y: -d.Y}
}
