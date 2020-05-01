package tree

import (
	"github.com/ivanterekh/qt-go-examples/internal/geometry"
	"sort"
)

type treeBuilder struct {
	barrier chan struct{}
}

func (tb *treeBuilder) newTree(points []geometry.Point, split split) *Tree {
	if len(points) == 0 {
		return nil
	}

	//tb.barrier <- struct{}{}
	//defer func() { <-tb.barrier }()

	t := &Tree{
		split: split,
	}

	if len(points) == 1 {
		t.point = points[0]
		return t
	}

	byX := func(i, j int) bool {
		return points[i].X < points[j].X
	}
	byY := func(i, j int) bool {
		return points[i].Y < points[j].Y
	}

	if split == verticalSplit {
		sort.Slice(points, byX)
	} else {
		sort.Slice(points, byY)
	}

	mid := len(points) / 2
	t.point = points[mid]

	//go func() {
	//	t.left = tb.newTree(points[:mid], split.opposite())
	//}()
	//go func() {
	//	t.right = tb.newTree(points[mid+1:], split.opposite())
	//}()
	t.left = tb.newTree(points[:mid], split.opposite())
	t.right = tb.newTree(points[mid+1:], split.opposite())

	return t
}
