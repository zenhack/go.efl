package elementary

// #include <Elementary.h>
import "C"

func (o *Object) NewButton() *Object {
	return (*Object)(C.elm_button_add((*C.Evas_Object)(o)))
}
