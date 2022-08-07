package tutorials

import (
	"fyne.io/fyne/v2"
)

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]Tutorial{
		"welcome":           {"欢迎", "", welcomeScreen, true},
		"timetool":          {"时间工具", "提供时间戳和日期互转", timetoolScreen, false},
		"md5tool":           {"md5工具", "提供生成md5的工具", md5toolScreen, false},
		"base64tool":        {"base64工具", "提供base编解码", base64toolScreen, false},
		"jsontool":          {"json操作", "提供json相关操作", jsontoolScreen, false},
		"json2gostructtool": {"json转go struct", "提供json快转go结构体", json2gostructtoolScreen, false},
		"terminaltool":      {"本地终端", "", terminaltoolScreen, false},
		// "imagetool": {"图片", "快速得到图片的base编码", imagetoolScreen, false},  //当文字太长时 组件NewMultiLineEntry 会引起崩溃，作者说会在2.3中修复
		"canvas": {"Canvas",
			"See the canvas capabilities.",
			canvasScreen,
			true,
		},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
		"": {"welcome", "timetool", "md5tool", "base64tool", "jsontool", "json2gostructtool", "terminaltool"},
		// "collections": {"list", "table", "tree"},//支持二级菜单
	}
)
