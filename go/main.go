package main

// #cgo CFLAGS: -I${SRCDIR}/../java/target
// #cgo LDFLAGS: -L${SRCDIR}/../java/target
// #cgo LDFLAGS: -lmyapp
// #include <libmyapp.h>
import "C"

import "fmt"

func main() {
	fmt.Println("Hi from go")
	C.run_main(0, nil)

	isolate := C.Java_dev_avelier_createIsolate()
	summator := C.Java_dev_avelier_createSummator(isolate)
	for i := 1; i < 10; i++ {
		x := C.Java_dev_avelier_callCalc(isolate, summator, C.double(1.0/float64(i)))
		fmt.Println(x)
	}
}
