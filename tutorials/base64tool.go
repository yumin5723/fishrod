package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"encoding/base64"
)

func base64toolScreen(_ fyne.Window) fyne.CanvasObject {
	currentArea := widget.NewMultiLineEntry()
	currentArea.Wrapping = fyne.TextWrapWord

	targetArea := widget.NewMultiLineEntry()
	targetArea.Wrapping = fyne.TextWrapWord

	buttonBox := container.NewVBox(widget.NewButton("编码", func() {
			str := currentArea.Text
			targetArea.SetText(base64.StdEncoding.EncodeToString([]byte(str)))
		}),
		widget.NewButton("解码", func() {
			str := currentArea.Text
			sDec, _ := base64.StdEncoding.DecodeString(str)
			targetArea.SetText(string(sDec))
		}),
		widget.NewButton("交换", func() {
			tmp := currentArea.Text
			currentArea.SetText(targetArea.Text)
			targetArea.SetText(tmp)
		}),
	)
	

	// content := container.NewVBox(
	// 	currentArea, buttonBox, targetArea,
	// )

	return container.NewGridWithRows(1, currentArea, buttonBox, targetArea)
}
