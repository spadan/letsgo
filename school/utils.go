// byte数组与string相互转换，零拷贝
package school

import (
	"reflect"
	"unsafe"
)

func String2BytesWithoutCopy(a *string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(a))
	sliceHeader := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}

func Bytes2StringWithoutCopy(a *[]byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(a))
	stringHeader := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&stringHeader))
}
