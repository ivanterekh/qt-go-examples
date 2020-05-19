package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"

	"github.com/ivanterekh/qt-go-examples/internal/color"
	"github.com/ivanterekh/qt-go-examples/internal/convexhull"
	"github.com/ivanterekh/qt-go-examples/internal/geometry"
	"github.com/ivanterekh/qt-go-examples/internal/window"
)

const margin = 10

const (
	jarvis = "Jarvis"
	quick  = "Quick"
	graham = "Graham"
)

type app struct {
	w      window.Window
	img    *gui.QImage
	points []geometry.Point

	getMethod func() string
}

func newApp(dx, dy int) app {
	a := app{
		w: window.New(dx, dy),
	}

	a.setupWidgets()
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

	genButton := widgets.NewQPushButton2("Generate", nil)
	genButton.ConnectClicked(func(bool) {
		a.inputHandler(input.Text())()
	})
	widget.Layout().AddWidget(genButton)

	methods := widgets.NewQComboBox(nil)
	methods.AddItems([]string{jarvis, quick, graham})
	a.getMethod = methods.CurrentText
	widget.Layout().AddWidget(methods)

	computeButton := widgets.NewQPushButton2("Compute", nil)
	computeButton.ConnectClicked(func(bool) {
		a.drawConvexHull()
	})
	widget.Layout().AddWidget(computeButton)

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

		a.genPoints(pointNum)
	}
}

func (a *app) genPoints(pointNum int) {
	a.img = a.w.Img()
	size := a.img.Size()
	w, h := size.Width(), size.Height()

	a.img.Fill2(color.Black)

	a.points = make([]geometry.Point, pointNum)

	for i := 0; i < pointNum; i++ {
		point := geometry.Point{
			X: rand.Intn(w-margin*2) + margin,
			Y: rand.Intn(h-margin*2) + margin,
		}
		a.points[i] = point
		a.img.SetPixelColor2(point.X, point.Y, color.Green)
	}

	a.w.SetImg(a.img)
}

func (a *app) drawConvexHull() {
	var hull []geometry.Point
	m := a.getMethod()
	switch m {
	case quick:
		hull = convexhull.SolveQuick(a.points)
	case jarvis:
		hull = convexhull.SolveJarvis(a.points)
	case graham:
		hull = convexhull.SolveGraham(a.points)
	default:
		log.Printf("unknown method %s", m)
		return
	}

	log.Printf("using %s for convex hull computation", m)

	painter := gui.NewQPainter2(a.img)
	defer painter.DestroyQPainter()

	pen := gui.NewQPen3(color.Red)
	pen.SetWidth(2)
	painter.SetPen(pen)

	for i := 0; i < len(hull)-1; i++ {
		drawLine(painter, hull[i], hull[i+1])
	}
	drawLine(painter, hull[0], hull[len(hull)-1])

	a.w.SetImg(a.img)
}

func drawLine(painter *gui.QPainter, p1, p2 geometry.Point) {
	painter.DrawLine3(p1.X, p1.Y, p2.X, p2.Y)
}
