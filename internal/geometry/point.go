package geometry

type Point struct {
	X, Y int
}

type PointF struct {
	X, Y float64
}

func (p1 PointF) add(p2 PointF) PointF {
	return PointF{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func FurthestPoint(points []Point, a, b Point) Point {
	var maxIdx int
	maxDst := dst(points[maxIdx], a, b)
	for i, p := range points {
		if maxDst < dst(p, a, b) {
			maxDst = dst(p, a, b)
			maxIdx = i
		}
	}
	return points[maxIdx]
}

// dst returns a square of distance between p and line connecting a and b
func dst(p, a, b Point) float64 {
	return float64(sq((b.Y-a.Y)*p.X-(b.X-a.X)*p.Y+b.X*a.Y-b.Y*a.X)) / float64(sq(b.Y-b.Y)+sq(b.X-a.X))
}

func sq(a int) int {
	return a * a
}
