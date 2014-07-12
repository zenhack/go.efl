package main
// This is a transliteration of the C program at:
//
//    http://docs.enlightenment.org/stable/elementary/group__Start.html


import (
	"unsafe"

	elm "github.com/zenhack/go.efl/elementary"
)

func main() {
	win := elm.NewStandardWindow("hello", "Hello")
	win.SmartCallbackAdd("delete,request", func(_ unsafe.Pointer) {
		elm.Exit()
	})

	box := win.NewBox()
	box.SetHorizontal(true)

	(*elm.Window)(win).ResizeAdd((*elm.Object)(box))
	box.Show()

	label := win.NewLabel()
	label.TextSet("Hello out there world!")
	label.Show()

	btn := win.NewButton()
	btn.TextSet("OK")
	box.PackEnd(btn)
	btn.Show()
	btn.SmartCallbackAdd("clicked", func(_ unsafe.Pointer) {
		elm.Exit()
	})
	win.Show()

	elm.Run()
}
