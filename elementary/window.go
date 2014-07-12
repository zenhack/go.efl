package elementary

// #include <Elementary.h>
import "C"

type Window Object

func (w *Window) ResizeAdd(o *Object) {
	C.elm_win_resize_object_add((*C.Evas_Object)(w), (*C.Evas_Object)(o))
}
