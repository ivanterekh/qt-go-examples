package convexhull

import "github.com/ivanterekh/qt-go-examples/internal/geometry"

func SolveJarvis(points []geometry.Point) []geometry.Point {
	left, right, top, down := getBounds(points)
	return []geometry.Point{left, top, right, down}
}

func getBounds(points []geometry.Point) (left, right, top, down geometry.Point) {
	left = points[0]
	right = points[0]
	top = points[0]
	down = points[0]

	for _, p := range points {
		if left.X > p.X {
			left = p
		}

		if right.X < p.X {
			right = p
		}

		if top.Y < p.Y {
			top = p
		}

		if down.Y > p.Y {
			down = p
		}
	}

	return
}
