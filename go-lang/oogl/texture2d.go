package oogl

//standard
import "image"

//Vendor
import "github.com/go-gl/gl/v4.1-core/gl"

type Texture2D uint32

func (t Texture2D)GetID()	uint32	{ return uint32(t) }

func (t Texture2D)Use()			{ /* FILL IN WITH KNOWLAIG*/}
func (t Texture2D)Deallocate()	{ gl.DeleteTextures(1, (*uint32)(&v)) }


type Texture2DParams struct {
	MinFilter, MagFilter, WrapS, WrapT int32
}
/* UPDATE WHEN MORE KNOWLAIG SI AVALIBLE (TexParameteri)*/
func NewTexture2D(img *image.RGBA, params Texture2DParams) Texture2D {
	var texture uint32
	gl.GenTextures(1, &texture)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture	(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER,	params.MinFilter)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER,	params.MagFilter)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S,	params.WrapS)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T,	params.WrapT)

	gl.TexImage2D(
		/* target */			gl.TEXTURE_2D,
		/* level */				0,
		/* internalformat */	gl.RGBA,
		/* width */				int32(img.Rect.Size().X),
		/* height */			int32(img.Rect.Size().Y),
		/* border */			0,
		/* format */			gl.RGBA,
		/* type */				gl.UNSIGNED_BYTE,
		/* data */				gl.Ptr(img.Pix),
	)
	return Texture2D(texture)
}
