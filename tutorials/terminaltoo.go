package tutorials

import (
	"fyne.io/fyne/v2"
	"github.com/fyne-io/terminal"
)

func terminaltoolScreen(win fyne.Window) fyne.CanvasObject {
	t := terminal.New()
	go func() {
		_ = t.RunLocalShell()
		// a.Quit()
	}()
	return t
}
