package window

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"

	"github.com/ivanterekh/qt-go-examples/internal/color"
)

type Window struct {
	*widgets.QMainWindow
	scene     *widgets.QGraphicsScene
	View      *widgets.QGraphicsView
	item      *widgets.QGraphicsPixmapItem
	Statusbar *widgets.QStatusBar
}

func (w Window) Img() *gui.QImage {
	return w.item.Pixmap().ToImage()
}

func (w Window) SetImg(img *gui.QImage) {
	w.item.SetPixmap(gui.NewQPixmap().FromImage(img, 0))
}

func New(dx, dy int) Window {
	var w Window
	w.QMainWindow = widgets.NewQMainWindow(nil, 0)
	w.QMainWindow.SetMinimumSize2(480, 720) // TODO: delete?

	w.Statusbar = widgets.NewQStatusBar(w.QMainWindow)
	w.QMainWindow.SetStatusBar(w.Statusbar)

	w.scene = widgets.NewQGraphicsScene(nil)
	w.View = widgets.NewQGraphicsView(nil)

	pixmap := gui.NewQPixmap2(w.View.Size())
	pixmap.Fill(color.Black)

	w.item = widgets.NewQGraphicsPixmapItem2(pixmap, nil)

	w.scene.AddItem(w.item)

	w.View.SetScene(w.scene)
	w.View.Show()

	w.View.ConnectResizeEvent(func(e *gui.QResizeEvent) {
		img := w.item.Pixmap().ToImage()
		w.SetImg(img.Scaled(e.Size(), core.Qt__IgnoreAspectRatio, core.Qt__FastTransformation))
	})

	w.QMainWindow.SetCentralWidget(w.View)
	w.QMainWindow.Show()

	return w
}
