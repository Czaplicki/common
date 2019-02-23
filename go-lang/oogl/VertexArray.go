package oogl

import 	"github.com/go-gl/gl/v4.1-core/gl"

type VertexArray uint32

func NewVertexArray() VertexArray {
	var id uint32
	gl.GenVertexArrays(1, &id)
	return VertexArray(id)
}

func (v VertexArray)GetID() uint32	{ return uint32(v)					}

func (v VertexArray)Use() 			{ gl.BindVertexArray(uint32(v))		}

func (v VertexArray)Deallocate()	{ gl.DeleteVertexArrays(1, (*uint32)(&v))	}
