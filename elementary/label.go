package elementary

// #include <Elementary.h>
import "C"

func (o *Object) NewLabel() *Object {
	return (*Object)(C.elm_label_add((*C.Evas_Object)(o)))
}
