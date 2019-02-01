package oogl

import 	"github.com/go-gl/gl/v4.1-core/gl"

type Vao uint32

func NewVao() Vao {
	var id uint32
	gl.GenVertexArrays(1, &id)
	return Vao(id)
}

func (v Vao)GetID() uint32	{ return uint32(v)					}

func (v Vao)Use() 			{ gl.BindVertexArray(uint32(v))		}

func (v Vao)Deallocate()	{ gl.DeleteVertexArrays(1, (*uint32)(&v))	}
