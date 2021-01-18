package main

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type tapIcon struct {
	widget.Icon
	tap func()
}

func newTapIcon(res fyne.Resource, fn func()) *tapIcon {
	i := &tapIcon{tap: fn}
	i.Resource = res
	i.ExtendBaseWidget(i)
	return i
}

func (t *tapIcon) Tapped(_ *fyne.PointEvent) {
	t.tap()
}

func makeUI() fyne.CanvasObject {
	return container.NewHBox(
		newTapIcon(theme.HomeIcon(), func() {
			log.Println("Go home")
		}),
		newTapIcon(theme.NavigateBackIcon(), func() {
			log.Println("Go back")
		}),
		newTapIcon(theme.NavigateNextIcon(), func() {
			log.Println("Go forward")
		}),
	)
}

func main() {
	a := app.New()
	w := a.NewWindow("Navigate")

	w.SetContent(makeUI())
	w.ShowAndRun()
}