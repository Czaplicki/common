package oogl


//Vendor
import "github.com/go-gl/gl/v4.1-core/gl"

import "github.com/Czaplicki/common/go-lang/texture"

type Texture2D uint32

func (t Texture2D)GetID()	uint32	{ return uint32(t) }

func (t Texture2D)Use(index uint32){
		gl.ActiveTexture(gl.TEXTURE0 + index)
		gl.BindTexture	(gl.TEXTURE_2D, uint32(t))
}
func (t Texture2D)Deallocate()	{ gl.DeleteTextures(1, (*uint32)(&t))		}

func (t Texture2D)SetOption(option uint32, value int32) {
	gl.BindTexture	(gl.TEXTURE_2D, uint32(t))
	gl.TexParameteri(gl.TEXTURE_2D, option, value)
}

func NewTexture2D() Texture2D {
	var texture uint32
	gl.GenTextures(1, &texture)
	return Texture2D(texture)
}
func (t Texture2D)Load(img *texture.Texture) {

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture	(gl.TEXTURE_2D, uint32(t))
	gl.TexImage2D(
		/* target */			gl.TEXTURE_2D,
		/* level */				0,
		/* internalformat */	gl.RGBA,
		/* width */				int32(img.Width),
		/* height */			int32(img.Height),
		/* border */			0,
		/* format */			gl.RGBA,
		/* type */				gl.UNSIGNED_BYTE,
		/* data */				gl.Ptr(img.Pixels),
	)
}
