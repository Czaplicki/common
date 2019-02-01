package oogl

import 	"github.com/go-gl/gl/v4.1-core/gl"

type ArrayBuffer uint32

func NewArrayBuffer() ArrayBuffer {
	var id uint32
	gl.GenBuffers(1, &id)
	return ArrayBuffer(id)
}
func NewArrayBuffers(n int) []ArrayBuffer {
	buffers := make([]ArrayBuffer, n, n)
	gl.GenBuffers(1, (*uint32)(&(buffers[0])))
	return buffers
}

func (b ArrayBuffer)GetID() uint32	{ return 			  uint32(b)		}
func (b ArrayBuffer)Use()			{ gl.BindBuffer		( gl.ARRAY_BUFFER	,    uint32(  b) )}
func (b ArrayBuffer)Deallocate()	{ gl.DeleteBuffers	( 1					,  (*uint32)(&b) )}

/* TODO: Add gl.usages to GPU pgk*/
func (b ArrayBuffer)FillFloat32(data []float32) {
	b.Use()
	gl.BufferData(	/*buffer type*/ gl.ARRAY_BUFFER,
		 			/*sizeInBytes*/ len(data) * 4 ,
					/*	Ptr		 */ gl.Ptr(data),
					/* usage	 */ gl.STATIC_DRAW,
				)
}
