package oogl

import "unsafe"
import	"github.com/go-gl/gl/v4.1-core/gl"

//base functions
var Init	func()				error	= gl.Init
var Clear	func(mask uint32)			= gl.Clear

//rendering
var DrawElements func(mode uint32, count int32, xtype uint32, indices unsafe.Pointer) = gl.DrawElements


//aaaaaaaa
var EnableVertexAttribArray func (index uint32) = gl.EnableVertexAttribArray
var GetAttribLocation func (program uint32, name *uint8) int32 = gl.GetAttribLocation
var VertexAttribPointer func (index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer) = gl.VertexAttribPointer

var BindFragDataLocation func (program uint32, color uint32, name *uint8) = gl.BindFragDataLocation

//Convertions
var Ptr			func (data interface{})	unsafe.Pointer	= gl.Ptr
var PtrOffset	func (offset	int)		 unsafe.Pointer	= gl.PtrOffset
var Str			func (str		string)		*uint8			= gl.Str
var GoString	func (cstr		*uint8)		 string			= gl.GoStr

var Strs		func (strs	 ...string)	(cstrs **uint8, free func())	= gl.Strs


//oogl utility

func Deallocate(resources ...GPUResource) {
	for _, resource := range resources {
		resource.Deallocate()
	}
}
