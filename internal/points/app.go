package main

import (
	"fmt"
	"github.com/ivanterekh/qt-go-examples/internal/color"
	"github.com/ivanterekh/qt-go-examples/internal/geometry"
	"github.com/ivanterekh/qt-go-examples/internal/tree"
	"github.com/ivanterekh/qt-go-examples/internal/window"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"math/rand"
	"strconv"
)

type app struct {
	w        window.Window
	dx, dy   int
	pointNum int
	tree     tree.Tree
}

func newApp(dx, dy int) app {
	a := app{
		w:        window.New(dx, dy),
		pointNum: 1000,
	}

	a.setupWidgets()
	return a
}

func (a app) setupWidgets() {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQHBoxLayout())
	widget.Resize2(widget.Width(), 40)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Number of points")
	widget.Layout().AddWidget(input)

	button := widgets.NewQPushButton2("Generate", nil)
	button.ConnectClicked(func(bool) {
		pointNum, err := strconv.Atoi(input.Text())
		if err != nil {
			widgets.QMessageBox_Information(
				nil,
				"Error",
				fmt.Sprintf("\"%s\" is not a number", input.Text()),
				widgets.QMessageBox__Ok,
				widgets.QMessageBox__Ok)
			return
		}

		a.pointNum = pointNum
		a.genPoints()
	})
	widget.Layout().AddWidget(button)

	dock := widgets.NewQDockWidget2(nil, 0)
	dock.SetWidget(widget)
	dock.SetFeatures(0)
	a.w.AddDockWidget(core.Qt__TopDockWidgetArea, dock)
}

func (a app) genPoints() {
	size := a.w.View.Size()
	w, h := size.Width(), size.Height()

	img := a.w.Img()
	img.Fill2(color.Black)

	points := make([]geometry.Point, a.pointNum)

	for i := 0; i < a.pointNum; i++ {
		point := geometry.Point{
			X: rand.Intn(w),
			Y: rand.Intn(h),
		}
		points[i] = point
		img.SetPixelColor2(point.X, point.Y, color.Green)
	}

	a.tree = tree.New(points)
	a.w.SetImg(img)
}
