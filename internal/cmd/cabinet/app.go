package main

import (
	"github.com/ivanterekh/qt-go-examples/internal/window"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type app struct {
	w window.Window
}

func newApp(dx, dy int) app {
	a := app{
		w: window.New(dx, dy),
	}

	a.w.ConnectPaintEvent(func(*gui.QPaintEvent) {

	})

	a.setupWidgets()
	a.connectEventHandlers()
	return a
}

func (a *app) setupWidgets() {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQHBoxLayout())
	widget.Resize2(widget.Width(), 40)

	x := widgets.NewQSlider(nil)
	x.SetWindowIconText("x angle")
	widget.Layout().AddWidget(x)


	dock := widgets.NewQDockWidget2(nil, 0)
	dock.SetWidget(widget)
	dock.SetFeatures(0)
	a.w.AddDockWidget(core.Qt__TopDockWidgetArea, dock)
}

func (a *app) connectEventHandlers() {
}
