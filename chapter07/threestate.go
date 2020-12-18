package main

import (
	"image/color"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type CheckState int

const (
	CheckOff CheckState = iota
	CheckOn
	CheckIndeterminate
)

type ThreeStateCheck struct {
	widget.BaseWidget
	State CheckState
}

func NewThreeStateCheck() *ThreeStateCheck {
	c := &ThreeStateCheck{}
	c.ExtendBaseWidget(c)
	return c
}

func (c *ThreeStateCheck) Tapped(_ *fyne.PointEvent) {
	if c.State == CheckIndeterminate {
		c.State = CheckOff
	} else {
		c.State++
	}

	c.Refresh()
}

func (c *ThreeStateCheck) CreateRenderer() fyne.WidgetRenderer {
	r := &threeStateRender{check: c, img: &canvas.Image{}}
	r.updateImage()
	return r
}

type threeStateRender struct {
	check *ThreeStateCheck
	img   *canvas.Image
}

func (t *threeStateRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (t *threeStateRender) Destroy() {
}

func (t *threeStateRender) Layout(_ fyne.Size) {
	t.img.Resize(t.MinSize())
}

func (t *threeStateRender) MinSize() fyne.Size {
	return fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize())
}

func (t *threeStateRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{t.img}
}

func (t *threeStateRender) Refresh() {
	t.updateImage()
}

func (t *threeStateRender) updateImage() {
	defer t.img.Refresh()

	switch t.check.State {
	case CheckOn:
		t.img.Resource = theme.CheckButtonCheckedIcon()
	case CheckIndeterminate:
		res, err := fyne.LoadResourceFromPath("indeterminate_check_box-24px.svg")
		if err != nil {
			log.Println("Failed to load indeterminate resource")
			return
		}

		t.img.Resource = theme.NewThemedResource(res, nil)
	default:
		t.img.Resource = theme.CheckButtonIcon()
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Three State")

	w.SetContent(NewThreeStateCheck())
	w.ShowAndRun()
}
