package cfunc

// #include "cfunc.h"
// #include <stdlib.h>
import "C"
import "unsafe"

import (
	"fmt"
)

//export Show
func Show(name, msg *C.char) {
	goname := C.GoString(name)
	gomsg := C.GoString(msg)
	fmt.Printf("[go]%s: %s\n", goname, gomsg)
}

func Say(name, msg string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))
	C.Say(cname, cmsg)
}
