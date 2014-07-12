#include <Evas.h>
#include <stdint.h>

/* This is a thin wrapper used to get around a restriction in cgo (namely that
 * you can't pass go functions to C). For details on the whole procedure, see
 * the comment in the body of Object.SmartCallbackAdd. */
void c_do_callback(void *data, Evas_Object *obj, void *event_info);
