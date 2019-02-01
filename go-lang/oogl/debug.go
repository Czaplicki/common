package oogl

import "io"
import "fmt"
import "unsafe"
import "github.com/go-gl/gl/v4.1-core/gl"

import "runtime/debug"

var sources map[uint32]string = map[uint32]string {

	gl.DEBUG_SOURCE_API					: "API",
	gl.DEBUG_SOURCE_WINDOW_SYSTEM		: "WINDOW SYSTEM",
	gl.DEBUG_SOURCE_SHADER_COMPILER		: "SHADER COMPILER",
	gl.DEBUG_SOURCE_THIRD_PARTY			: "THIRD PARTY",
	gl.DEBUG_SOURCE_APPLICATION			: "APPLICATION",
	gl.DEBUG_SOURCE_OTHER				: "UNKNOWN",
}
var types map[uint32]string = map[uint32]string {

	gl.DEBUG_TYPE_ERROR					: "ERROR",
	gl.DEBUG_TYPE_DEPRECATED_BEHAVIOR	: "DEPRECATED BEHAVIOR",
	gl.DEBUG_TYPE_UNDEFINED_BEHAVIOR	: "UDEFINED BEHAVIOR",
	gl.DEBUG_TYPE_PORTABILITY			: "PORTABILITY",
	gl.DEBUG_TYPE_PERFORMANCE			: "PERFORMANCE",
	gl.DEBUG_TYPE_OTHER					: "OTHER",
	gl.DEBUG_TYPE_MARKER				: "MARKER",
}
var severitys map[uint32]string = map[uint32]string {

	gl.DEBUG_SEVERITY_HIGH				: "HIGH",
	gl.DEBUG_SEVERITY_MEDIUM			: "MEDIUM",
	gl.DEBUG_SEVERITY_LOW				: "LOW",
	gl.DEBUG_SEVERITY_NOTIFICATION		: "NOTIFICATION",
}

var output io.Writer
func AttachDebugProc(wr io.Writer) {

	output = wr
	gl.Enable(gl.DEBUG_OUTPUT)
	gl.DebugMessageCallback(debugProc, gl.Ptr(nil))
}
func debugProc(	source	uint32	, gltype	uint32,
				id		uint32	, severity	uint32,
				length	int32	, message	string,
				userParam	unsafe.Pointer		) {

	output.Write([]byte(fmt.Sprintf("%d: %s of %s severity, raised from %s: %s\n",
							id, types[gltype],
							severitys[severity],
							sources[source],
							 message)))
	//output.Write([]byte("\u001b[30m;1"))
	output.Write([]byte(debug.Stack()))
	//output.Write([]byte("\u001b[0m"))
	fmt.Println([]byte(fmt.Sprintf("%d: %s of %s severity, raised from %s: %s\n",
							id, types[gltype],
							severitys[severity],
							sources[source],
							 message)))
}
// !!!!!!!!!!!!!!!!!!!!!!!!!!!! READ how to set glDebugCallBack(func)!!!!!!!!!!!!!!!!!!!!!!!!
