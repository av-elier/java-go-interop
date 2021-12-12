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
}
