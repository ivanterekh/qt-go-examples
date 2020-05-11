package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	a := newApp(1500, 1000)

	widgets.QApplication_SetStyle2("fusion")
	a.w.Show()
	widgets.QApplication_Exec()
}
