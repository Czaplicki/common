package oogl

import	"github.com/go-gl/gl/v4.1-core/gl"

const (
	COLOR_BUFFER_BIT = gl.COLOR_BUFFER_BIT
	DEPTH_BUFFER_BIT = gl.DEPTH_BUFFER_BIT
)

func Clear(mask uint32) { gl.Clear(mask) }
