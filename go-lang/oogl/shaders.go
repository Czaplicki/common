package oogl

import "fmt"
import "strings"

import	"github.com/go-gl/gl/v4.1-core/gl"

import	"github.com/Czaplicki/common/go-lang/log"

type Shader uint32

const (
	GEOMETRY_SHADER			 uint32 = gl.GEOMETRY_SHADER
	VERTEX_SHADER			 uint32 = gl.VERTEX_SHADER
	FRAGMENT_SHADER			 uint32 = gl.FRAGMENT_SHADER
	COMPUTE_SHADER			 uint32 = gl.COMPUTE_SHADER
	TESS_CONTROL_SHADER		 uint32 = gl.TESS_CONTROL_SHADER
	TESS_EVALUATION_SHADER	 uint32 = gl.TESS_EVALUATION_SHADER
)

func NewShader(source string, shaderType uint32) (Shader, error) {

	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	gl.CompileShader(shader)
	free()

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return Shader(shader), nil
}

func (s Shader)GetID() uint32	{ return uint32(s) }
func (s Shader)Use()			{ log.Error("Shader #", s, " does nothing on Use") }
func (s Shader)Deallocate()		{
	gl.DeleteShader(uint32(s))
}
func DeallocateShaders(shaders ...Shader) {
	for _, s :=  range shaders {
		s.Deallocate()
	}
}
