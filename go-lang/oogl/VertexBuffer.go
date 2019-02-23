package oogl

import "unsafe"

import 	"github.com/go-gl/gl/v4.1-core/gl"

type VertexBuffer struct {
	id, _type uint32

}
func NewVertexBuffer(_type uint32) VertexBuffer {
	buffer := VertexBuffer{}
	buffer._type = _type
	gl.GenBuffers(1, &buffer.id)
	return buffer
}

//GPUResource
func (b VertexBuffer)GetID() uint32		{ return b.id	}
func (b VertexBuffer)Use()				{ gl.BindBuffer		( b._type	,    b.id )}
func (b VertexBuffer)Deallocate()		{ gl.DeleteBuffers	( 1			,   &b.id )}
//VertexBuffer
func (b VertexBuffer)GetType() uint32	{ return b._type	}

/* TODO: Add gl.usages to oogl pgk*/
func (b VertexBuffer)Load(data unsafe.Pointer, bytelength int32, usage uint32) {
	b.Use()
	gl.BufferData(	/*buffer type*/ b._type,
		 			/*sizeInBytes*/ int(bytelength),
					/*	Ptr		 */ data,
					/* usage	 */ usage,
				)
}

//buffer types
const (
	ARRAY_BUFFER			uint32 = gl.ARRAY_BUFFER
	ELEMENT_ARRAY_BUFFER	uint32 = gl.ELEMENT_ARRAY_BUFFER
)

//buffer usages
const (
	STATIC_DRAW = gl.STATIC_DRAW
)
