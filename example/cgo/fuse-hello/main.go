package main

// #cgo CFLAGS: -D_FILE_OFFSET_BITS=64 -I/usr/include/fuse  -pthread
// #cgo LDFLAGS: -lfuse
//
// #include "helloMount.h"
import "C"
import "flag"
import "log"

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n  hello [MOUNTPOINT]")
	}

	mountPoint := flag.Arg(0)
	if mountPoint == "" {
		log.Fatal("Usage:\n  hello [MOUNTPOINT]")
	}
	C.helloMount(C.CString(mountPoint))
}
