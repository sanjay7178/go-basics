// package main

// import (
// 	"time"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )

// type Device struct {
// 	Name      string
// 	Status    string
// 	Timestamp string
// }

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Device List")

// 	devices := []Device{
// 		{Name: "Dev 1", Status: "Denied", Timestamp: "xx:xx"},
// 		{Name: "Dev 1", Status: "Allowed", Timestamp: "xx:xx"},
// 		{Name: "Dev 1", Status: "-", Timestamp: "xx:xx"},
// 	}

// 	table := widget.NewTable(
// 		func() (int, int) { return len(devices), 4 },
// 		func() fyne.CanvasObject {
// 			return container.NewHBox(
// 				widget.NewLabel("wide content"),
// 				widget.NewLabel("status"),
// 				widget.NewLabel("timestamp"),
// 				container.NewHBox(
// 					widget.NewButton("Allow", func() {}),
// 					widget.NewButton("Deny", func() {}),
// 				),
// 			)
// 		},
// 		func(id widget.TableCellID, cell fyne.CanvasObject) {
// 			container := cell.(*fyne.Container)
// 			device := devices[id.Row]
// 			switch id.Col {
// 			case 0:
// 				container.Objects[0].(*widget.Label).SetText(device.Name)
// 			case 1:
// 				container.Objects[1].(*widget.Label).SetText(device.Status)
// 			case 2:
// 				container.Objects[2].(*widget.Label).SetText(device.Timestamp)
// 			case 3:
// 				buttons := container.Objects[3].(*fyne.Container)
// 				allowBtn := buttons.Objects[0].(*widget.Button)
// 				denyBtn := buttons.Objects[1].(*widget.Button)
				
// 				if device.Status == "-" {
// 					allowBtn.Show()
// 					denyBtn.Show()
// 				} else {
// 					allowBtn.Hide()
// 					denyBtn.Hide()
// 					container.Objects[3].(*fyne.Container).Objects[0].(*widget.Label).SetText("-")
// 				}

// 				allowBtn.OnTapped = func() {
// 					devices[id.Row].Status = "Allowed"
// 					devices[id.Row].Timestamp = time.Now().Format("15:04")
// 					table.Refresh()
// 				}
// 				denyBtn.OnTapped = func() {
// 					devices[id.Row].Status = "Denied"
// 					devices[id.Row].Timestamp = time.Now().Format("15:04")
// 					table.Refresh()
// 				}
// 			}
// 		},
// 	)

// 	myWindow.SetContent(table)
// 	myWindow.Resize(fyne.NewSize(600, 300))
// 	myWindow.ShowAndRun()
// }


package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Device struct {
	Name      string
	Status    string
	Timestamp string
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Device List")

	devices := []Device{
		{Name: "Dev 1", Status: "Denied", Timestamp: "xx:xx"},
		{Name: "Dev 1", Status: "Allowed", Timestamp: "xx:xx"},
		{Name: "Dev 1", Status: "-", Timestamp: "xx:xx"},
	}

	var deviceTable *widget.Table
	deviceTable = widget.NewTable(
		func() (int, int) { return len(devices), 4 },
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("wide content"),
				widget.NewLabel("status"),
				widget.NewLabel("timestamp"),
				container.NewHBox(
					widget.NewButton("Allow", func() {}),
					widget.NewButton("Deny", func() {}),
				),
			)
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			container := cell.(*fyne.Container)
			device := devices[id.Row]
			switch id.Col {
			case 0:
				container.Objects[0].(*widget.Label).SetText(device.Name)
			case 1:
				container.Objects[1].(*widget.Label).SetText(device.Status)
			case 2:
				container.Objects[2].(*widget.Label).SetText(device.Timestamp)
			case 3:
				buttons := container.Objects[3].(*fyne.Container)
				allowBtn := buttons.Objects[0].(*widget.Button)
				denyBtn := buttons.Objects[1].(*widget.Button)
				
				if device.Status == "-" {
					allowBtn.Show()
					denyBtn.Show()
				} else {
					allowBtn.Hide()
					denyBtn.Hide()
					container.Objects[3].(*fyne.Container).Objects[0].(*widget.Label).SetText("-")
				}

				allowBtn.OnTapped = func() {
					devices[id.Row].Status = "Allowed"
					devices[id.Row].Timestamp = time.Now().Format("15:04")
					deviceTable.Refresh()
				}
				denyBtn.OnTapped = func() {
					devices[id.Row].Status = "Denied"
					devices[id.Row].Timestamp = time.Now().Format("15:04")
					deviceTable.Refresh()
				}
			}
		},
	)

	myWindow.SetContent(deviceTable)
	myWindow.Resize(fyne.NewSize(600, 300))
	myWindow.ShowAndRun()
}