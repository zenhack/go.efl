package elementary

// #include <Elementary.h>
import "C"

import (
	"os"
	"unsafe"
)

func init() {
	// We call elm_init here -- it's silly to have the user call it
	// manually, so we don't bother exposing it.
	argc := len(os.Args)
	argv := make([]*C.char, argc)
	for i := range argv {
		argv[i] = C.CString(os.Args[i])
		defer C.free(unsafe.Pointer(argv[i]))
	}
	C.elm_init(C.int(argc), &argv[0])
}

// Run starts elementary's main loop. It returns when Exit is called, after
// cleaning up any resources used by elementary.
func Run() {
	C.elm_run()
	C.elm_shutdown()
}

// Exits elementary's main loop.
func Exit() {
	C.elm_exit()
}
