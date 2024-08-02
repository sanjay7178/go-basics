package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

var data = [][]string{
    {"top left", "top right"},
    {"bottom left", "bottom right"},
}

func test() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Table Widget")

    list := widget.NewTable(
        func() (int, int) {
            return len(data), len(data[0])
        },
        func() fyne.CanvasObject {
            return container.NewHBox(
                widget.NewLabel(""),
                widget.NewButton("Allow", func() {}),
                widget.NewButton("Block", func() {}),
            )
        },
        func(i widget.TableCellID, o fyne.CanvasObject) {
            // Extract the label from the container and set its text
            container := o.(*fyne.Container)
            label := container.Objects[0].(*widget.Label)
            label.SetText(data[i.Row][i.Col])
        })

    myWindow.SetContent(list)
    myWindow.Resize(fyne.NewSize(400, 300))
    myWindow.ShowAndRun()
}