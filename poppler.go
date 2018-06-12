package poppler

// #cgo pkg-config: poppler-glib
// #include <poppler.h>
// #include <stdlib.h>
// #include <glib.h>
// #include <unistd.h>
import "C"

import (
	"errors"
	"path/filepath"
)

type poppDoc *C.struct__PopplerDocument

func Open(filename string) (doc *Document, err error) {
	filename, err = filepath.Abs(filename)
	if err != nil {
		return
	}
	var e *C.GError
	fn := C.g_filename_to_uri((*C.gchar)(C.CString(filename)), nil, nil)
	var d poppDoc
	d = C.poppler_document_new_from_file((*C.char)(fn), nil, &e)
	if e != nil {
		err = errors.New(C.GoString((*C.char)(e.message)))
	}
	doc = &Document{
		doc: d,
	}
	return
}

func Version() string {
	return C.GoString(C.poppler_get_version())
}
