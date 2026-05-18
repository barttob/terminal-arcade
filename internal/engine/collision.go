package engine

func IsCollision(a, b Point) bool {
	return a.X == b.X && a.Y == b.Y
}

func ContainsPoint(points []Point, target Point) bool {
	for _, p := range points {
		if p.X == target.X && p.Y == target.Y {
			return true
		}
	}
	return false
}
