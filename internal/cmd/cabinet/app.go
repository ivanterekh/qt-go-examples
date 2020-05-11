package main

import (
	"fmt"
	"math"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"

	"github.com/ivanterekh/qt-go-examples/internal/color"
	"github.com/ivanterekh/qt-go-examples/internal/projection"
	"github.com/ivanterekh/qt-go-examples/internal/window"
)

type app struct {
	w         window.Window
	axisAngle map[string]int
}

func newApp(dx, dy int) app {
	a := app{
		w:         window.New(dx, dy),
		axisAngle: make(map[string]int),
	}

	a.w.ConnectPaintEvent(func(*gui.QPaintEvent) {
		img := a.w.Img()

		imgSize := img.Size()
		cubeSize := math.Min(float64(imgSize.Height()), float64(imgSize.Width())) / 2
		center := gui.NewQVector3D3(float32(imgSize.Width()/2), float32(imgSize.Height()/2), 0)

		lines := projection.GetCubeProjection(center, cubeSize, a.axisAngle["x"], a.axisAngle["y"], a.axisAngle["z"])

		img.Fill2(color.Black)

		painter := gui.NewQPainter2(img)
		defer painter.DestroyQPainter()

		pen := gui.NewQPen3(color.Green)
		pen.SetWidth(3)
		painter.SetPen(pen)

		for _, line := range lines {
			painter.DrawLine(line)
		}

		a.w.SetImg(img)
	})

	a.setupWidgets()
	a.connectEventHandlers()
	return a
}

func (a *app) setupWidgets() {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQHBoxLayout())
	widget.Resize2(widget.Width(), 40)

	a.addSlider(widget, "x")
	a.addSlider(widget, "y")
	a.addSlider(widget, "z")

	dock := widgets.NewQDockWidget2(nil, 0)
	dock.SetWidget(widget)
	dock.SetFeatures(0)
	a.w.AddDockWidget(core.Qt__TopDockWidgetArea, dock)
}

func (a *app) addSlider(widget *widgets.QWidget, name string) {
	labelFmt := "%s=%d°"

	label := widgets.NewQLabel2(fmt.Sprintf(labelFmt, name, 0), nil, 0)
	label.SetFixedWidth(70)

	slider := widgets.NewQSlider(nil)
	slider.SetRange(0, 360)
	slider.ConnectSliderMoved(func(value int) {
		label.SetText(fmt.Sprintf(labelFmt, name, value))
		a.axisAngle[name] = value
		a.w.Update()
	})

	widget.Layout().AddWidget(label)
	widget.Layout().AddWidget(slider)
}

func (a *app) connectEventHandlers() {
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
