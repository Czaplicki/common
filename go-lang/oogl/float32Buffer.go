package oogl

import 	"github.com/go-gl/gl/v4.1-core/gl"

type Float32Buffer uint32

func NewFloat32Buffer() Float32Buffer {
	var id uint32
	gl.GenBuffers(1, &id)
	return ArrayBuffer(id)
}

func (b ArrayBuffer)GetID() uint32	{ return 			  uint32(b)		}
func (b ArrayBuffer)Use()			{ gl.BindBuffer		( gl.ARRAY_BUFFER	,    uint32(  b) )}
func (b ArrayBuffer)Deallocate()	{ gl.DeleteBuffers	( 1					,  (*uint32)(&b) )}

/* TODO: Add gl.usages to GPU pgk*/
func (b ArrayBuffer)Load(data []float32) {
	b.Use()
	gl.BufferData(	/*buffer type*/ gl.ARRAY_BUFFER,
		 			/*sizeInBytes*/ len(data) * 4 ,
					/*	Ptr		 */ gl.Ptr(data),
					/* usage	 */ gl.STATIC_DRAW,
				)
}
