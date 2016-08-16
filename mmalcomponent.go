package main

/*
#include <mmal.h>
*/
import "C"
import (
	"unsafe"
)

type (
	MMALComponentT C.struct_MMAL_COMPONENT_T
	MMMALSuccess   C.enum_MMAL_SUCCESS
)

func main() {
	// Create the video decoder component on VideoCore
	decoder := &MMALComponentT{}
	component := C.CString("vc.ril.video_decoder")
	status := mmal_component_create(component, &decoder)
	C.ABORT_IF_ERROR(status)

	// Set the format of the video decoder input port
	formatIn := *decoder.input

	defer C.free(unsafe.Pointer(decoder))
	defer C.free(unsafe.Pointer(component))
}
