package tutorials

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func jsontoolScreen(_ fyne.Window) fyne.CanvasObject {
	currentArea := widget.NewMultiLineEntry()
	currentArea.Wrapping = fyne.TextWrapWord

	targetArea := widget.NewMultiLineEntry()
	targetArea.Wrapping = fyne.TextWrapWord

	buttonBox := container.NewVBox(widget.NewButton("格式化", func() {
		data := currentArea.Text
		var str bytes.Buffer
		_ = json.Indent(&str, []byte(data), "", "    ")
		targetArea.SetText(str.String())
	}),
		widget.NewButton("压缩", func() {
			data := currentArea.Text
			var str bytes.Buffer
			json.Compact(&str, []byte(data))
			targetArea.SetText(str.String())
		}),
		widget.NewButton("转义", func() {
			targetArea.SetText(strconv.Quote(currentArea.Text))
		}),
		widget.NewButton("去除转义", func() {
			data := currentArea.Text
			data = strings.ReplaceAll(data, "\\\"", "\"")
			data = strings.ReplaceAll(data, "\\\\", "\\")
			// s, err := strconv.Unquote(data)
			// fmt.Println(err.Error())

			targetArea.SetText(data)
		}),
		widget.NewButton("Unicode转中文", func() {
			data := currentArea.Text
			s, _ := strconv.Unquote(strings.Replace(strconv.Quote(data), `\\u`, `\u`, -1))
			targetArea.SetText(s)
		}),
		widget.NewButton("中文转Unicode", func() {
			data := currentArea.Text
			textQuoted := strconv.QuoteToASCII(data)
			textUnquoted := textQuoted[1 : len(textQuoted)-1]
			textUnquoted = strings.ReplaceAll(textUnquoted, "\\\"", "\"")
			textUnquoted = strings.ReplaceAll(textUnquoted, "\\\\", "\\")
			// fmt.Println(textUnquoted)
			// s, err := strconv.Unquote(textUnquoted)
			// fmt.Println(err.Error())
			targetArea.SetText(textUnquoted)
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
