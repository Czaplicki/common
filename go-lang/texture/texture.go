package texture

import "image"
import "image/color"


func NewPixel(c color.Color) (Pixel){
	r, g, b, a := c.RGBA()
	return Pixel{
		R : byte(r),
		G : byte(g),
		B : byte(b),
		A : byte(a),
	}
}
type Pixel struct {
	R, G, B, A byte
}
func (p Pixel)RGBA() (uint32, uint32, uint32, uint32) {
	return uint32(p.R), uint32(p.G), uint32(p.B), uint32(p.A)
}

type Texture struct {
	Width, Height int
	Pixels []Pixel
}

func NewTexture(img image.Image) Texture {
	imgbounds := img.Bounds()
	texture := Texture{
		Width	: imgbounds.Max.X - imgbounds.Min.X,
		Height	: imgbounds.Max.Y - imgbounds.Min.Y,
	}
	texture.Pixels = make([]Pixel, texture.Width * texture.Height)

	for i := 0; i < len(texture.Pixels); i++ {
		x := i % texture.Width + imgbounds.Min.X
		y := i / texture.Width + imgbounds.Min.X
		texture.Pixels[i] = NewPixel(img.At(x, y))
	}
	return texture
}

func (t Texture)ColorModel() color.Model {
	return color.RGBAModel
}

func (t Texture)Bounds() image.Rectangle {
	return image.Rect(0, 0, t.Width, t.Height)
}

func (t Texture)At(x, y int) color.Color {
	return t.Pixels[y * t.Width + x]
}
