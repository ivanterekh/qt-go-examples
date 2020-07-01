package convexhull

import (
	"math"
	"sort"

	"github.com/ivanterekh/qt-go-examples/internal/geometry"
)

func SolveGraham(points []geometry.Point) []geometry.Point {
	start := points[0]
	for _, p := range points {
		if p.X < start.X || (p.X == start.X && p.Y < start.Y) {
			start = p
		}
	}
	sort.Slice(points, func(i, j int) bool {
		tgi, tgj := tg(start, points[i]), tg(start, points[j])
		if tgi == tgj {
			return points[i].Y > points[j].Y
		}
		return tgi < tgj
	})

	res := []geometry.Point{start, points[1]}
	for _, p := range points {
		last, prev := getLast2(res)
		if p == last || p == prev {
			continue
		}

		prd := prod(last, prev, p)
		for prd < 0 || (prd == 0 && p.X > last.X && last.X < prev.X) {
			res = res[:len(res)-1]
			last, prev = getLast2(res)
			prd = prod(last, prev, p)
		}

		res = append(res, p)
	}

	return res
}

func tg(a, b geometry.Point) float64 {
	if a.X == b.X {
		if a.Y > b.Y {
			return math.MaxFloat64
		}
		return -math.MaxFloat64
	}
	return float64(a.Y-b.Y) / float64(a.X-b.X)
}

func getLast2(pts []geometry.Point) (geometry.Point, geometry.Point) {
	return pts[len(pts)-1], pts[len(pts)-2]
}
