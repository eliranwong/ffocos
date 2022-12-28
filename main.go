package main

import (
	"os/exec"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("ffocos")

	// message for display
	message := widget.NewRichTextFromMarkdown(`
# ffocos
## fix fcitx on Chrome OS

To keep 'fcitx' working with fcitx-partially-compatible apps on Chrome OS,

simply keep this window opened or minized while you are using gui apps.

For information, read:

[https://github.com/eliranwong/ffocos/blob/main/README.md](https://github.com/eliranwong/ffocos/blob/main/README.md)`)

	w.SetContent(message)
	w.ShowAndRun()

	exec.Command("killall", "-9", "ffocos")
}
