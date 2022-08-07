# fishrod
一个类似“站长工具”的桌面应用

# Getting Started
将main.go中init方法设置字体替换为自己本地字体，以支持中文输出，如：
```
func init() {
	os.Setenv("FYNE_FONT", "/System/Library/Fonts/STHeiti Medium.ttc")

}
```
And you can run that simply as:

$ go run main.go
