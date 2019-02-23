package oogl

import	"github.com/go-gl/gl/v4.1-core/gl"

func (u Uniform)SetFloat1(v0 float32) {
	gl.Uniform1f(int32(u), v0)
}


func (u Uniform)SetFloat1Array(count int32, value *float32) {
	gl.Uniform1fv(int32(u), count, value)
}


func (u Uniform)SetInt1(v0 int32) {
	gl.Uniform1i(int32(u), v0)
}


func (u Uniform)SetInt1Array(count int32, value *int32) {
	gl.Uniform1iv(int32(u), count, value)
}


func (u Uniform)SetUInt1(v0 uint32) {
	gl.Uniform1ui(int32(u), v0)
}

func (u Uniform)SetUInt1Array(count int32, value *uint32) {
	gl.Uniform1uiv(int32(u), count, value)
}

func (u Uniform)SetDecimal1(x float64) {
	gl.Uniform1d(int32(u), x)
}
func (u Uniform)SetDecimal1Array(count int32, value *float64) {
	gl.Uniform1dv(int32(u), count, value)
}



func (u Uniform)SetFloat2(v0 float32, v1 float32) {
	gl.Uniform2f(int32(u), v0, v1)
}


func (u Uniform)SetFloat2Array(count int32, value *float32) {
	gl.Uniform2fv(int32(u), count, value)
}


func (u Uniform)SetInt2(v0 int32, v1 int32) {
	gl.Uniform2i(int32(u), v0, v1)
}

func (u Uniform)SetInt2Array(count int32, value *int32) {
	gl.Uniform2iv(int32(u), count, value)
}


func (u Uniform)SetUInt2(v0 uint32, v1 uint32) {
	gl.Uniform2ui(int32(u), v0, v1)
}

func (u Uniform)SetUInt2Array(count int32, value *uint32) {
	gl.Uniform2uiv(int32(u), count, value)
}

func (u Uniform)SetDecimal2(x float64, y float64) {
	gl.Uniform2d(int32(u), x, y)
}
func (u Uniform)SetDecimal2Array(count int32, value *float64) {
	gl.Uniform2dv(int32(u), count, value)
}



func (u Uniform)SetFloat3(v0 float32, v1 float32, v2 float32) {
	gl.Uniform3f(int32(u), v0, v1, v2)
}


func (u Uniform)SetFloat3Array(count int32, value *float32) {
	gl.Uniform3fv(int32(u), count, value)
}


func (u Uniform)SetInt3(v0 int32, v1 int32, v2 int32) {
	gl.Uniform3i(int32(u), v0, v1, v2)
}

func (u Uniform)SetInt3Array(count int32, value *int32) {
	gl.Uniform3iv(int32(u), count, value)
}


func (u Uniform)SetUInt3(v0 uint32, v1 uint32, v2 uint32) {
	gl.Uniform3ui(int32(u), v0, v1, v2)
}

func (u Uniform)SetUInt3Array(count int32, value *uint32) {
	gl.Uniform3uiv(int32(u), count, value)
}

func (u Uniform)SetDecimal3(x float64, y float64, z float64) {
	gl.Uniform3d(int32(u), x, y, z)
}
func (u Uniform)SetDecimal3Array(count int32, value *float64) {
	gl.Uniform3dv(int32(u), count, value)
}

func (u Uniform)SetFloat4(v0 float32, v1 float32, v2 float32, v3 float32) {
	gl.Uniform4f(int32(u), v0, v1, v2, v3)
}


func (u Uniform)SetFloat4Array(count int32, value *float32) {
	gl.Uniform4fv(int32(u), count, value)
}


func (u Uniform)SetInt4(v0 int32, v1 int32, v2 int32, v3 int32) {
	gl.Uniform4i(int32(u), v0, v1, v2, v3)
}

func (u Uniform)SetInt4Array(count int32, value *int32) {
	gl.Uniform4iv(int32(u), count, value)
}


func (u Uniform)SetUInt4(v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	gl.Uniform4ui(int32(u), v0, v1, v2, v3)
}

func (u Uniform)SetUInt4Array(count int32, value *uint32) {
	gl.Uniform4uiv(int32(u), count, value)
}

func (u Uniform)SetDecimal4(x float64, y float64, z float64, w float64) {
	gl.Uniform4d(int32(u), x, y, z, w)
}
func (u Uniform)SetDecimal4Array(count int32, value *float64) {
	gl.Uniform4dv(int32(u), count, value)
}
