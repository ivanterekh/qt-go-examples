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
