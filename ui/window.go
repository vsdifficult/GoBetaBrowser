package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/zserge/lorca"
)

func CreateBrowserWindow(app fyne.App) fyne.Window {
	myWindow := app.NewWindow("Mini Browser")

	urlEntry := widget.NewEntry()
	urlEntry.SetText("https://www.google.com")

	goButton := widget.NewButton("Go", func() {
		openLorca(urlEntry.Text)
	})

	myWindow.SetContent(container.NewVBox(
		urlEntry,
		goButton,
	))

	myWindow.Resize(fyne.NewSize(400, 200))

	openLorca(urlEntry.Text)

	return myWindow
}

func openLorca(url string) {
	ui, err := lorca.New(url, "", 800, 600)
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	<-ui.Done()
}
