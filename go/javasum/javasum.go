package javasum

// #cgo CFLAGS: -I${SRCDIR}/../../java/target
// #cgo LDFLAGS: -L${SRCDIR}/../../java/target
// #cgo LDFLAGS: -lmyapp
// #include <libmyapp.h>
import "C"

import (
	"unsafe"
)

type JavaSum struct {
	isolate  C.longlong
	summator unsafe.Pointer
}

func CreateJavaSum() JavaSum {
	isolate := C.Java_dev_avelier_createIsolate()
	summator := C.Java_dev_avelier_createSummator(isolate)
	return JavaSum{
		isolate:  isolate,
		summator: summator,
	}
}

func (s JavaSum) Calc(x float64) float64 {
	return float64(C.Java_dev_avelier_callCalc(s.isolate, s.summator, C.double(x)))
}
