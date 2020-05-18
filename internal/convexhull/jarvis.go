package convexhull

import (
	"github.com/ivanterekh/qt-go-examples/internal/geometry"
)

func SolveJarvis(points []geometry.Point) []geometry.Point {
	var res []geometry.Point

	left, right, top, down := getBounds(points)

	appendSector(
		&res,
		left, top,
		points,
		func(curr, next geometry.Point) float64 {
			return float64(next.Y-curr.Y) / float64(next.X-curr.X)
		},
		func(curr, next geometry.Point) bool {
			return curr != next && curr.X <= next.X && curr.Y <= next.Y
		},
		func(curr, cand geometry.Point) bool {
			return curr.X == cand.X
		},
	)

	appendSector(
		&res,
		top, right,
		points,
		func(curr, next geometry.Point) float64 {
			return float64(next.X-curr.X) / float64(curr.Y-next.Y)
		},
		func(curr, next geometry.Point) bool {
			return curr != next && curr.X <= next.X && curr.Y >= next.Y
		},
		func(curr, cand geometry.Point) bool {
			return curr.Y == cand.Y
		},
	)

	appendSector(
		&res,
		right, down,
		points,
		func(curr, next geometry.Point) float64 {
			return float64(curr.Y-next.Y) / float64(curr.X-next.X)
		},
		func(curr, next geometry.Point) bool {
			return curr != next && curr.X >= next.X && curr.Y >= next.Y
		},
		func(curr, cand geometry.Point) bool {
			return curr.X == cand.X
		},
	)

	appendSector(
		&res,
		down, left,
		points,
		func(curr, next geometry.Point) float64 {
			return float64(curr.X-next.X) / float64(next.Y-curr.Y)
		},
		func(curr, next geometry.Point) bool {
			return curr != next && curr.X >= next.X && curr.Y <= next.Y
		},
		func(curr, cand geometry.Point) bool {
			return curr.Y == cand.Y
		},
	)

	return res
}

func appendSector(
	res *[]geometry.Point,
	start, end geometry.Point,
	points []geometry.Point,
	getTg func(curr, next geometry.Point) float64,
	check func(curr, next geometry.Point) bool,
	isNext func(curr, cand geometry.Point) bool,
) {
	curr := start
	for curr != end {
		var tg float64
		var cand = end
		for _, p := range points {
			if !check(curr, p) {
				continue
			}

			if isNext(curr, p) {
				cand = p
				break
			}

			tgp := getTg(curr, p)
			if tg < tgp {
				tg = tgp
				cand = p
			}
		}

		*res = append(*res, cand)
		curr = cand
	}
}
