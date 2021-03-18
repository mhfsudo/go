package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func loadUI() fyne.CanvasObject {
	//summerList := []string{"Badehose", "Badetuch"}
	//winterList := []string{"Ski", "Skischuhe"}

	content := widget.NewMultiLineEntry()
	content.SetText("Content")

	// list
	list := widget.NewButton("Button1", func() {

		})

	// sidebar
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {

		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {

		}),
	)

	// contentbar

	side := container.New(layout.NewBorderLayout(bar, nil, nil, nil),bar, list)

	split := container.NewHSplit(side, content)

	return split

}

/*
func loadUI() fyne.CanvasObject {
content := widget.NewMultiLineEntry()

   list := widget.NewVBox(
      widget.NewButton("Inhalt 1", func () {
         content.SetText("Inhalt 1")
      }),
      widget.NewButton("Inhalt 2", func () {
         content.SetText("Inhalt 2")
      }),
   )
   bar := widget.NewToolbar(
      widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
      widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
	)

   side := fyne.NewContainerWithLayout(layout.NewBorderLayout(bar, nil, nil, nil), bar, list)
   split:= widget.NewHSplitContainer(side, content)
   split.Offset = 0.25
   return split
}
*/

func main() {
   a := app.New()
   w:= a.NewWindow("Packliste")
   //w.SetContent(widget.NewLabel("Hello"))
   w.SetContent(loadUI())
   w.ShowAndRun()
}