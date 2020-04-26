package window

import (
	"github.com/ivanterekh/qt-go-examples/internal/color"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Window struct {
	*widgets.QMainWindow
	scene     *widgets.QGraphicsScene
	view      *widgets.QGraphicsView
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
	w.QMainWindow.SetMinimumSize2(360, 520) // TODO: delete?

	w.Statusbar = widgets.NewQStatusBar(w.QMainWindow)
	w.QMainWindow.SetStatusBar(w.Statusbar)

	w.scene = widgets.NewQGraphicsScene(nil)
	w.view = widgets.NewQGraphicsView(nil)

	pixmap := gui.NewQPixmap2(core.NewQSize2(dx, dy))
	pixmap.Fill(color.Black)

	w.item = widgets.NewQGraphicsPixmapItem2(pixmap, nil)

	w.scene.AddItem(w.item)
	w.view.SetScene(w.scene)
	w.view.Show()

	w.QMainWindow.SetCentralWidget(w.view)
	w.QMainWindow.Show()

	return w
}
