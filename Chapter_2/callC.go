package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C function!")
	C.cHello()

	fmt.Println("Going to call another C function!")
	myMessage := C.CString("This is Mihalis!")
	/*нам понадобится оператор defer, чтобы освободить
	память, занимаемую строкой C, после того как она будет использована.
	Оператор defer включает в себя вызов функции C.free() и затем unsafe.Pointer().*/
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)

	fmt.Println("All perfectly done!")
}
