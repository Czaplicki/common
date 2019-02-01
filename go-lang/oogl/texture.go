package oogl

//standard
import "image"

//Vendor
import "github.com/go-gl/gl/v4.1-core/gl"

type Texture2D uint32

func (t Texture2D)GetID()	uint32	{ return uint32(t) }

func (t Texture2D)Use()			{ /* FILL IN WITH KNOWLAIG*/}
func (t Texture2D)Deallocate()	{ /* FILL IN WITH KNOWLAIG*/}


type Texture2DParams struct {
	MinFilter, MagFilter, WrapS, WrapT int32
}
/* UPDATE WHEN MORE KNOWLAIG SI AVALIBLE (TexParameteri)*/
func NewTexture(img *image.RGBA, params Texture2DParams) Texture2D {
	var texture uint32
	gl.GenTextures(1, &texture)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture	(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER,	params.MinFilter)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER,	params.MagFilter)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S,	params.WrapS)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T,	params.WrapT)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(img.Rect.Size().X),
		int32(img.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(img.Pix),
	)
	return Texture2D(texture)
}
