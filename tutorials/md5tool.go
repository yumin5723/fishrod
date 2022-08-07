package tutorials

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func md5toolScreen(_ fyne.Window) fyne.CanvasObject {
	md5Area := widget.NewMultiLineEntry()
	md5Area.Wrapping = fyne.TextWrapWord

	md5Result32Upper := widget.NewEntry()
	md5Result32Upper.SetPlaceHolder("32位大")

	md5Result32Lower := widget.NewEntry()
	md5Result32Lower.SetPlaceHolder("32位小")

	md5Result16Upper := widget.NewEntry()
	md5Result16Upper.SetPlaceHolder("16位大")

	md5Result16Lower := widget.NewEntry()
	md5Result16Lower.SetPlaceHolder("16位小")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "32位大", Widget: md5Result32Upper, HintText: ""},
			{Text: "32位小", Widget: md5Result32Lower, HintText: ""},
			{Text: "16位大", Widget: md5Result16Upper, HintText: ""},
			{Text: "16位小", Widget: md5Result16Lower, HintText: ""},
		},
	}

	buttonBox := container.NewHBox(widget.NewButton("加密", func() {
		md5Str := md5Area.Text
		if md5Str != "" {
			hash := md5.Sum([]byte(md5Str))
			md5Result32Upper.SetText(strings.ToUpper(hex.EncodeToString(hash[:])))
			md5Result32Lower.SetText(strings.ToLower(hex.EncodeToString(hash[:])))
			md5Result16Upper.SetText(strings.ToUpper(hex.EncodeToString(hash[:]))[8:24])
			md5Result16Lower.SetText(strings.ToLower(hex.EncodeToString(hash[:]))[8:24])
		}
	}),
		widget.NewButton("清空", func() {
			md5Area.SetText("")
			md5Result32Upper.SetText("")
			md5Result32Lower.SetText("")
			md5Result16Upper.SetText("")
			md5Result16Lower.SetText("")
			
		}),
	)

	content := container.NewVBox(
		md5Area, buttonBox, form,
	)

	return container.NewGridWithRows(1, content)
	// return container.NewGridWithRows(2, content, content1)
	// return
	// return container.NewVBox(content, content1)
}
