package elementary

// #include <Elementary.h>
import "C"

type Box Object

func (o *Object) NewBox() *Box {
	return (*Box)(C.elm_box_add((*C.Evas_Object)(o)))
}

func (b *Box) SetHorizontal(horizontal bool) {
	var value C.Eina_Bool
	if horizontal {
		value = C.EINA_TRUE
	} else {
		value = C.EINA_FALSE
	}
	C.elm_box_horizontal_set((*C.Evas_Object)(b), value)
}

func (b *Box) Show() {
	(*Object)(b).Show()
}

func (b *Box) PackEnd(o *Object) {
	C.elm_box_pack_end((*C.Evas_Object)(b), (*C.Evas_Object)(o))
}
