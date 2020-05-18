package convexhull

import (
	"github.com/ivanterekh/qt-go-examples/internal/geometry"
)

func SolveQuick(points []geometry.Point) []geometry.Point {
	left, right, _, _ := getBounds(points)
	u := getSector(points, left, right, upper)
	l := getSector(points, left, right, lower)

	u = append(u, right)
	for i := len(l) - 1; i >= 0; i-- {
		u = append(u, l[i])
	}
	u = append(u, left)

	return u
}

func getSector(
	points []geometry.Point, left, right geometry.Point,
	onSide func(p, l, r geometry.Point) bool,
) []geometry.Point {
	points = filter(points, func(p geometry.Point) bool {
		return onSide(p, left, right)
	})

	if len(points) < 2 {
		return points
	}

	f := geometry.FurthestPoint(points, left, right)
	res := getSector(points, left, f, onSide)
	res = append(res, f)
	return append(res, getSector(points, f, right, onSide)...)
}

func filter(points []geometry.Point, cond func(p geometry.Point) bool) []geometry.Point {
	var res []geometry.Point
	for _, p := range points {
		if cond(p) {
			res = append(res, p)
		}
	}
	return res
}

func upper(p, l, r geometry.Point) bool {
	return prod(p, l, r) > 0
}

func lower(p, l, r geometry.Point) bool {
	return prod(p, l, r) < 0
}

func prod(p, l, r geometry.Point) int {
	return (p.X-l.X)*(r.Y-l.Y) - (r.X-l.X)*(p.Y-l.Y)
}
