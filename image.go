package poppler

// #cgo pkg-config: poppler-glib
// #include <poppler.h>
// #include <glib.h>
import "C"
import "unsafe"
import "github.com/ungerik/go-cairo"

// Image

type Image struct {
	Id   int
	Area Rectangle
	p    *C.struct__PopplerPage
}

func (im *Image) GetSurface() (cs *cairo.Surface) {
	ci := C.poppler_page_get_image(im.p, C.gint(im.Id))
	ctx := C.cairo_create(ci)
	cip := (cairo.Cairo_surface)(unsafe.Pointer(ci))
	ctxp := (cairo.Cairo_context)(unsafe.Pointer(ctx))
	return cairo.NewSurfaceFromC(cip, ctxp)
}
