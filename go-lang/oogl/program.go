package oogl

import "fmt"
import "strings"

import	"github.com/go-gl/gl/v4.1-core/gl"

type Program uint32

func (p Program)GetID() uint32 	{ return 			uint32(p)	}
func (p Program)Use()			{ gl.UseProgram(	uint32(p) )	}
func (p Program)Deallocate()	{ gl.DeleteProgram(	uint32(p) )	}


func (p Program)BindFragDataLocation(colorIndex uint32, name string) {
	glName	:= name + "\x00"
	gl.BindFragDataLocation(uint32(p), colorIndex, gl.Str(glName))
}

func (p Program)GetAttributeLocation(name string) (uint32, error) {
	glName	:= name + "\x00"
	index	:= gl.GetAttribLocation(p.GetID(), gl.Str(glName))
	if index < 0 { return 0, fmt.Errorf("Program: %v, doesn't contain any attribute named : %s", p, name) }
	return uint32(index), nil
}
//func (p Program)SetAttribute(name string, _type AttributeType, AttributeLayout)

func (p Program)GetUniform(name string) (Uniform, error) {
	glName := name + "\x00"
	ul := gl.GetUniformLocation(uint32(p), gl.Str(glName))
	var err error
	if ul > -1 {
		err = fmt.Errorf("Program: %v, doesn't contain any uniform named : \"%s\" (%x)", p, name, glName)
	}
	return Uniform(ul), err
}

func NewEmptyProgram() (Program) {
	return Program(gl.CreateProgram())
}

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

		gl.DeleteProgram(p)
		return 0, fmt.Errorf("Failed to link program: %v", log)
	}

	for _, s := range shaders {
		gl.DetachShader(p, uint32(s))
	}

	return Program(p), nil
}

func (p Program)Link(shaders ...Shader) error {

	id := uint32(p)
	for _, s := range shaders {
		gl.AttachShader(id, uint32(s))
	}
	gl.LinkProgram(id)

	var status int32
	gl.GetProgramiv(id, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(id, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(id, logLength, nil, gl.Str(log))

		return fmt.Errorf("failed to link program: %v", log)
	}

	for _, s := range shaders {
		gl.DetachShader(id, uint32(s))
	}

	return nil
}
