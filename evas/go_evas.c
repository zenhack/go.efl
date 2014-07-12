#include <stdint.h>
#include "go_evas.h"
#include "_cgo_export.h"

void c_do_callback(void *data, Evas_Object *obj, void *event_info) {
	goDoCallback((uintptr_t)data, event_info);
}
