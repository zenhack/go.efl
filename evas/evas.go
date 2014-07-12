package evas

// #cgo pkg-config: evas
// #include <Evas.h>
// #include "go_evas.h"
import "C"

import (
	"unsafe"
)

type Object C.Evas_Object

func (o *Object) Show() {
	C.evas_object_show((*C.Evas_Object)(o))
}

// Wraps evas_smart_callback_add. The data argument is ommitted; instead, the caller
// should wrap any extra data needed in a closure.
func (o *Object) SmartCallbackAdd(event string, callback func(evInfo unsafe.Pointer)) {
	// cgo doesn't allow us to pass a go function to C, so as a work around,
	// we store the function in a map (callbacks), and pass the key as the data
	// argument. c_do_callback is just a thin wrapper which invokes goDoCallback,
	// below, which calls the actual handler.
	obj := (*C.Evas_Object)(o)
	cStr := C.CString(event) // FIXME: memory leak
	ref := nextCallback
	nextCallback++
	callbacks[ref] = callback
	C.evas_object_smart_callback_add(obj, cStr, (*[0]byte)(C.c_do_callback),
		unsafe.Pointer(ref))
}

var (
	callbacks    map[uintptr]func(event_info unsafe.Pointer)
	nextCallback uintptr
)

//export goDoCallback
func goDoCallback(ref uintptr, event_info unsafe.Pointer) {
	callbacks[ref](event_info)
}

func init() {
	callbacks = make(map[uintptr]func(event_info unsafe.Pointer))
}
