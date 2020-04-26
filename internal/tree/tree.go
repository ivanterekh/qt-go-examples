package tree

import (
	"github.com/ivanterekh/qt-go-examples/internal/geometry"
	"github.com/therecipe/qt/core"
	"runtime"
)

type split int

const (
	verticalSplit split = iota + 1
	horizontalSplit
)

func (s split) opposite() split {
	if s == verticalSplit {
		return horizontalSplit
	}
	return verticalSplit
}

type Tree struct {
	point       geometry.Point
	split       split
	left, right *Tree
}

func New(points []geometry.Point) Tree {
	tb := treeBuilder{
		barrier: make(chan struct{}, runtime.NumCPU()-2),
	}

	return *tb.newTree(points, verticalSplit)
}

func (t Tree) intersect(rect core.QRect) bool {
	return (t.split == verticalSplit && t.point.X >= rect.Left() && t.point.X <= rect.Right()) ||
		(t.split == horizontalSplit && t.point.Y >= rect.Top() && t.point.Y <= rect.Bottom())
}

func (t Tree) pointInside(rect core.QRect) bool {
	return rect.Contains3(t.point.X, t.point.Y)
}

func (t Tree) rectInLeft(rect core.QRect) bool {
	return (t.split == verticalSplit && t.point.X > rect.Right()) ||
		(t.split == horizontalSplit && t.point.Y > rect.Bottom())
}

func (t Tree) PointsInRect(rect core.QRect) []geometry.Point {
	var res []geometry.Point

	if t.intersect(rect) {
		if t.pointInside(rect) {
			res = append(res, t.point)
		}
		insertNodePoints(t.left, rect, &res)
		insertNodePoints(t.right, rect, &res)
	} else if t.rectInLeft(rect) {
		insertNodePoints(t.left, rect, &res)
	} else {
		insertNodePoints(t.left, rect, &res)
	}

	return res
}

func insertNodePoints(tree *Tree, rect core.QRect, res *[]geometry.Point) {
	if tree != nil {
		*res = append(*res, tree.PointsInRect(rect)...)
	}
}
