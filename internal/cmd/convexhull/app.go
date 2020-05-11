package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"github.com/ivanterekh/qt-go-examples/internal/window"
)

type app struct {
	w         window.Window
}

func newApp(dx, dy int) app {
	a := app{
		w:         window.New(dx, dy),
	}

	a.setupWidgets()
	return a
}

func (a *app) setupWidgets() {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQHBoxLayout())
	widget.Resize2(widget.Width(), 40)

	dock := widgets.NewQDockWidget2(nil, 0)
	dock.SetWidget(widget)
	dock.SetFeatures(0)
	a.w.AddDockWidget(core.Qt__TopDockWidgetArea, dock)
}
