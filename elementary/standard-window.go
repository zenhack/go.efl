package elementary

// #include <Elementary.h>
import "C"

import (
	"unsafe"
)

func NewStandardWindow(name, title string) *Object {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	obj := C.elm_win_util_standard_add(cName, cTitle)
	return (*Object)(obj)
}
