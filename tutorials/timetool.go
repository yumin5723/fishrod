package tutorials

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
)

func timetoolScreen(_ fyne.Window) fyne.CanvasObject {
	current := time.Now()
	timestamp := strconv.FormatInt(current.Unix(), 10)
	text1 := widget.NewLabel("当前日期:" + current.Format("2006-01-02 15:04:05"))
	text2 := widget.NewLabel("当前时间戳(秒):" + strconv.FormatInt(current.Unix(), 10))
	copyButton := widget.NewButton("复制当前时间戳", func() {
		clipboard.WriteAll(timestamp)
	})

	label := container.NewHBox(text1, text2, copyButton)

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Unix时间戳")

	unit := widget.NewSelect([]string{"秒", "毫秒"}, func(s string) { fmt.Println("selected", s) })
	unit.SetSelected("秒")

	result := widget.NewEntry()
	result.SetPlaceHolder("结果")

	content := container.NewVBox(
		label,
		entry, unit, widget.NewButton("转换", func() {
			if unit.Selected == "秒" {
				unixTime, err := strconv.Atoi(entry.Text)
				if err != nil {
					result.SetText("时间戳错误")
				} else {
					tm := time.Unix(int64(unixTime), 0)
					result.SetText(tm.Format("2006-01-02 15:04:05"))
				}
			} else {
				msInt, err := strconv.ParseInt(entry.Text, 10, 64)
				if err != nil {
					result.SetText("时间戳错误")
				} else {
					tm := time.Unix(0, msInt*int64(time.Millisecond))
					result.SetText(tm.Format("2006-02-01 15:04:05"))
				}
			}

		}), result,
	)

	year := widget.NewEntry()
	year.SetPlaceHolder("年：如2022")
	fmt.Println(year.MinSize())

	month := widget.NewEntry()
	month.SetPlaceHolder("月：如08")

	day := widget.NewEntry()
	day.SetPlaceHolder("日：如31")

	hour := widget.NewEntry()
	hour.SetPlaceHolder("时：如23")

	min := widget.NewEntry()
	min.SetPlaceHolder("分：如59")

	second := widget.NewEntry()
	second.SetPlaceHolder("秒：如59")

	unit1 := widget.NewSelect([]string{"秒", "毫秒"}, func(s string) { fmt.Println("selected", s) })
	unit1.SetSelected("秒")

	result1 := widget.NewEntry()
	result1.SetPlaceHolder("结果")

	content1 := container.NewVBox(
		year, month, day, hour, min, second, unit1, widget.NewButton("转换", func() {
			timeStr := strings.TrimSpace(year.Text) + "-" + strings.TrimSpace(month.Text) + "-" + strings.TrimSpace(day.Text) + " " + strings.TrimSpace(hour.Text) + ":" + strings.TrimSpace(min.Text) + ":" + strings.TrimSpace(second.Text)
			fmt.Println(timeStr)
			times, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
			if unit1.Selected == "秒" {
				result1.SetText(strconv.Itoa(int(times.Unix())))
			} else {
				result1.SetText(strconv.Itoa(int(1000 * times.Unix())))
			}
		}), result1,
	)
	return container.NewGridWithRows(1, content, content1)
	// return container.NewGridWithRows(2, content, content1)
	// return
	// return container.NewVBox(content, content1)
}
