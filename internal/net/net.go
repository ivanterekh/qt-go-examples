package net

import (
	"math"

	"github.com/therecipe/qt/core"

	"github.com/ivanterekh/qt-go-examples/internal/geometry"
)

type Net struct {
	net    [][][]geometry.Point
	cw, ch int
}

func New(points []geometry.Point, w, h int) *Net {
	n := len(points)
	k := int(math.Ceil(math.Pow(float64(n), 0.33)))

	net := make([][][]geometry.Point, k)
	for i := 0; i < k; i++ {
		net[i] = make([][]geometry.Point, k)
	}

	cw, ch := w/k+1, h/k+1
	for _, p := range points {
		cell := &net[p.Y/ch][p.X/cw]
		*cell = append(*cell, p)
	}

	return &Net{
		net: net,
		cw:  cw,
		ch:  ch,
	}
}

func (pn Net) PointsInRect(rect core.QRect) []geometry.Point {
	var res []geometry.Point

	iMin := rect.Y() / pn.ch
	iMax := (rect.Y() + rect.Height()) / pn.ch
	jMin := rect.X() / pn.cw
	jMax := (rect.X() + rect.Width()) / pn.cw

	if iMin < 0 {
		iMin = 0
	}
	if jMin < 0 {
		jMin = 0
	}
	if iMax >= len(pn.net) {
		iMax = len(pn.net) - 1
	}
	if jMax >= len(pn.net[0]) {
		jMax = len(pn.net[0]) - 1
	}

	for i := iMin; i <= iMax; i++ {
		for j := jMin; j <= jMax; j++ {
			if i == iMin || i == iMax || j == jMin || j == jMax {
				for _, p := range pn.net[i][j] {
					if rect.Contains(core.NewQPoint2(p.X, p.Y), true) {
						res = append(res, p)
					}
				}
			} else {
				res = append(res, pn.net[i][j]...)
			}
		}
	}

	return res
}
