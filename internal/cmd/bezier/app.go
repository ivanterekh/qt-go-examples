package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"

	"github.com/ivanterekh/qt-go-examples/internal/color"
	"github.com/ivanterekh/qt-go-examples/internal/geometry"
	"github.com/ivanterekh/qt-go-examples/internal/window"
)

const circleR = 7

type app struct {
	w   window.Window

	points     []*geometry.Point
	movedPoint *geometry.Point
}

func newApp(dx, dy int) app {
	a := app{
		w: window.New(dx, dy),
	}

	a.connectEventHandlers()
	return a
}

func (a *app) connectEventHandlers() {
	a.w.View.ConnectMousePressEvent(a.mousePressHandler)
	a.w.View.ConnectMouseMoveEvent(a.mouseMoveHandler)
	a.w.ConnectPaintEvent(a.paintHandler)
}

func (a *app) mouseMoveHandler(e *gui.QMouseEvent) {
	a.movedPoint.X = e.X()
	a.movedPoint.Y = e.Y()

	a.w.Update()
}

func (a *app) mousePressHandler(e *gui.QMouseEvent) {
	a.movedPoint = a.getPoint(e.Pos())

	if a.movedPoint == nil {
		a.movedPoint = &geometry.Point{
			X: e.X(),
			Y: e.Y(),
		}
		a.points = append(a.points, a.movedPoint)
	}

	a.w.Update()
}

func (a *app) getPoint(pos *core.QPoint) *geometry.Point {
	for _, p := range a.points {
		if absInt(p.X - pos.X()) < circleR && absInt(p.Y - pos.Y()) < circleR {
			return p
		}
	}
	return nil
}

func (a *app) paintHandler(event *gui.QPaintEvent) {
	img := gui.NewQImage10(a.w.Img())
	defer a.w.SetImg(img)

	img.Fill2(color.Black)

	painter := gui.NewQPainter2(img)
	defer painter.DestroyQPainter()

	painter.SetPen(gui.NewQPen3(color.Green))
	painter.SetBrush(gui.NewQBrush3(color.Green, 0))

	for _, p := range a.points {
		painter.DrawEllipse5(core.NewQPoint2(p.X, p.Y), circleR, circleR)
	}
}

func absInt(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
