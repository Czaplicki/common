package byteutil

import "C"
import "unsafe"

func GetBytesInt32(value int32) []byte {
	return C.GoBytes(unsafe.Pointer(&value), 4)
}
func GetBytesUInt32(value uint32) []byte {
	return C.GoBytes(unsafe.Pointer(&value), 4)
}
func GetBytesFloat32(value float32) []byte {
	return C.GoBytes(unsafe.Pointer(&value), 4)
}
func GetBytesFloat64(value float32) []byte {
	return C.GoBytes(unsafe.Pointer(&value), 4)
}
