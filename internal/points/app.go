package main

import (
	"fmt"
	"github.com/ivanterekh/qt-go-examples/internal/color"
	"github.com/ivanterekh/qt-go-examples/internal/geometry"
	"github.com/ivanterekh/qt-go-examples/internal/tree"
	"github.com/ivanterekh/qt-go-examples/internal/window"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"math/rand"
	"strconv"
)

type app struct {
	w         window.Window
	img       *gui.QImage
	tree      *tree.Tree
	pointNum  int
	rectDrawn bool

	rectStart, rectEnd *core.QPoint
}

func newApp(dx, dy int) app {
	a := app{
		w:        window.New(dx, dy),
		pointNum: 1000,
	}

	a.img = a.w.Img()
	a.setupWidgets()
	a.connectEventHandlers()
	return a
}

func (a *app) setupWidgets() {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQHBoxLayout())
	widget.Resize2(widget.Width(), 40)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Number of points")
	input.ConnectReturnPressed(a.inputHandler(input.Text()))
	widget.Layout().AddWidget(input)

	button := widgets.NewQPushButton2("Generate", nil)
	button.ConnectClicked(func(bool) {
		a.inputHandler(input.Text())()
	})
	widget.Layout().AddWidget(button)

	dock := widgets.NewQDockWidget2(nil, 0)
	dock.SetWidget(widget)
	dock.SetFeatures(0)
	a.w.AddDockWidget(core.Qt__TopDockWidgetArea, dock)
}

func (a *app) inputHandler(text string) func() {
	return func() {
		if len(text) == 0 {
			return
		}

		pointNum, err := strconv.Atoi(text)
		if err != nil {
			widgets.QMessageBox_Information(
				nil,
				"Error",
				fmt.Sprintf("\"%s\" is not a number", text),
				widgets.QMessageBox__Ok,
				widgets.QMessageBox__Ok)
			return
		}

		a.pointNum = pointNum
		a.genPoints()
	}
}

func (a *app) genPoints() {
	a.img = a.w.Img()
	size := a.img.Size()
	w, h := size.Width(), size.Height()

	a.img.Fill2(color.Black)

	points := make([]geometry.Point, a.pointNum)

	for i := 0; i < a.pointNum; i++ {
		point := geometry.Point{
			X: rand.Intn(w),
			Y: rand.Intn(h),
		}
		points[i] = point
		a.img.SetPixelColor2(point.X, point.Y, color.Green)
	}

	a.tree = tree.New(points)
	a.rectStart, a.rectEnd = nil, nil
	a.w.SetImg(a.img)
}

func (a *app) connectEventHandlers() {
	a.w.View.ConnectMousePressEvent(a.mousePressHandler)
	a.w.View.ConnectMouseMoveEvent(a.mouseMoveHandler)
	a.w.ConnectPaintEvent(a.paintHandler)
}

func (a *app) mouseMoveHandler(event *gui.QMouseEvent) {
	a.rectEnd = event.Pos()
	a.w.Update()
}

func (a *app) mousePressHandler(event *gui.QMouseEvent) {
	if a.rectDrawn {
		// TODO
	} else {
		a.rectStart = event.Pos()
		a.rectEnd = event.Pos()
	}
	a.w.Update()
}

func (a *app) paintHandler(event *gui.QPaintEvent) {
	img := gui.NewQImage10(a.img)

	if a.tree == nil {
		return
	}

	rect := a.rect()
	for _, point := range a.tree.PointsInRect(*rect) {
		img.SetPixelColor2(point.X, point.Y, color.Red)
	}

	painter := gui.NewQPainter2(img)
	painter.SetPen(gui.NewQPen3(color.White))
	painter.DrawRect(core.NewQRectF5(rect))

	a.w.SetImg(img)
}

func (a *app) rect() *core.QRect {
	if a.rectStart == nil || a.rectEnd == nil {
		return core.NewQRect()
	}

	var (
		x1 = a.rectStart.X()
		x2 = a.rectEnd.X()
		y1 = a.rectStart.Y()
		y2 = a.rectEnd.Y()
	)

	if x1 > x2 {
		x1, x2 = x2, x1
	}

	if y1 > y2 {
		y1, y2 = y2, y1
	}

	return core.NewQRect2(core.NewQPoint2(x1, y1), core.NewQPoint2(x2, y2))
}
