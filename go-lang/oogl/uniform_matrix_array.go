package oogl

import	"github.com/go-gl/gl/v4.1-core/gl"

func (u Uniform)SetMatrix2dArray(count int32, transpose bool, value *float64) {
	gl.UniformMatrix2dv(int32(u), count, transpose, value)
}


func (u Uniform)SetMatrix2fArray(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2fv(int32(u), count, transpose, value)
}
func (u Uniform)SetMatrix2x3dArray(count int32, transpose bool, value *float64) {
	gl.UniformMatrix2x3dv(int32(u), count, transpose, value)
}


func (u Uniform)SetMatrix2x3fArray(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2x3fv(int32(u), count, transpose, value)
}
func (u Uniform)SetMatrix2x4dArray(count int32, transpose bool, value *float64) {
	gl.UniformMatrix2x4dv(int32(u), count, transpose, value)
}


func (u Uniform)SetMatrix2x4fArray(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2x4fv(int32(u), count, transpose, value)
}
func (u Uniform)SetMatrix3dArray(count int32, transpose bool, value *float64) {
	gl.UniformMatrix3dv(int32(u), count, transpose, value)
}


func (u Uniform)SetMatrix3fArray(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3fv(int32(u), count, transpose, value)
}
func (u Uniform)SetMatrix3x2dArray(count int32, transpose bool, value *float64) {
	gl.UniformMatrix3x2dv(int32(u), count, transpose, value)
}


func (u Uniform)SetMatrix3x2fArray(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3x2fv(int32(u), count, transpose, value)
}
func (u Uniform)SetMatrix3x4dArray(count int32, transpose bool, value *float64) {
	gl.UniformMatrix3x4dv(int32(u), count, transpose, value)
}


func (u Uniform)SetMatrix3x4fArray(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3x4fv(int32(u), count, transpose, value)
}
func (u Uniform)SetMatrix4dArray(count int32, transpose bool, value *float64) {
	gl.UniformMatrix4dv(int32(u), count, transpose, value)
}


func (u Uniform)SetMatrix4fArray(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4fv(int32(u), count, transpose, value)
}
func (u Uniform)SetMatrix4x2dArray(count int32, transpose bool, value *float64) {
	gl.UniformMatrix4x2dv(int32(u), count, transpose, value)
}


func (u Uniform)SetMatrix4x2fArray(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4x2fv(int32(u), count, transpose, value)
}
func (u Uniform)SetMatrix4x3dArray(count int32, transpose bool, value *float64) {
	gl.UniformMatrix4x3dv(int32(u), count, transpose, value)
}


func (u Uniform)SetMatrix4x3fArray(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4x3fv(int32(u), count, transpose, value)
}
