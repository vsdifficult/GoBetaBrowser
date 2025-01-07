package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/vsdifficult/GoBetaBrowser/ui"
)

func main() {
	myApp := app.New()
	myWindow := ui.CreateBrowserWindow(myApp)
	myWindow.ShowAndRun()
}
