package oogl

import "fmt"
import "strings"

import	"github.com/go-gl/gl/v4.1-core/gl"

type Program uint32

func (p Program)GetID() uint32 	{ return 			uint32(p)	}
func (p Program)Use()			{ gl.UseProgram(	uint32(p) )	}
func (p Program)Deallocate()	{ gl.DeleteProgram(	uint32(p) )	}

func NewProgram(shaders ...Shader) (Program, error) {

	p := gl.CreateProgram()

	for _, s := range shaders {
		gl.AttachShader(p, uint32(s))
	}
	gl.LinkProgram(p)

	var status int32
	gl.GetProgramiv(p, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(p, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(p, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	return Program(p), nil
}
