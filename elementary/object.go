package elementary

// #include <Elementary.h>
import "C"

import (
	"unsafe"

	"github.com/zenhack/go.efl/evas"
)

type Object evas.Object

func (o *Object) SmartCallbackAdd(event string, callback func(evInfo unsafe.Pointer)) {
	(*evas.Object)(o).SmartCallbackAdd(event, callback)
}

func (o *Object) Show() {
	(*evas.Object)(o).Show()
}

func (o *Object) TextSet(text string) {
	cText := C.CString(text)
	// XXX: elm_object_text set is defined as a macro, and for whatever reason cgo
	// isn't picking it up (getting an error to the effect that the symbol isn't
	// defined). instead, we're just calling the underlying function ourselves,
	// but it would be nice to not need to.
	//
	// C.elm_object_text_set((*C.Evas_Object)(o), cText)
	C.elm_object_part_text_set((*C.Evas_Object)(o), nil, cText)
}
