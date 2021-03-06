package oogl

import "io"
import "fmt"
import "unsafe"
import "github.com/go-gl/gl/v4.1-core/gl"

import "github.com/Czaplicki/common/go-lang/console/color"

import "runtime/debug"

const (
	RED = "\x00"
)

var sources map[uint32]string = map[uint32]string {

	gl.DEBUG_SOURCE_API					: "API",
	gl.DEBUG_SOURCE_WINDOW_SYSTEM		: "WINDOW SYSTEM",
	gl.DEBUG_SOURCE_SHADER_COMPILER		: "SHADER COMPILER",
	gl.DEBUG_SOURCE_THIRD_PARTY			: "THIRD PARTY",
	gl.DEBUG_SOURCE_APPLICATION			: "APPLICATION",
	gl.DEBUG_SOURCE_OTHER				: "UNKNOWN",
}
var types map[uint32]string = map[uint32]string {

	gl.DEBUG_TYPE_ERROR					: color.Red			( "ERROR"				),
	gl.DEBUG_TYPE_DEPRECATED_BEHAVIOR	: color.Magenta		( "DEPRECATED BEHAVIOR"	),
	gl.DEBUG_TYPE_UNDEFINED_BEHAVIOR	: color.BrightRed	( "UDEFINED BEHAVIOR"	),
	gl.DEBUG_TYPE_PORTABILITY			: color.Cyan		( "PORTABILITY"			),
	gl.DEBUG_TYPE_PERFORMANCE			: color.Blue		( "PERFORMANCE"			),
	gl.DEBUG_TYPE_OTHER					: color.Yellow		( "OTHER"				),
	gl.DEBUG_TYPE_MARKER				: color.BrightBlue	( "MARKER"				),
}
var severitys map[uint32]string = map[uint32]string {

	gl.DEBUG_SEVERITY_HIGH				: color.Red			( "HIGH"			),
	gl.DEBUG_SEVERITY_MEDIUM			: color.BrightRed	( "MEDIUM"			),
	gl.DEBUG_SEVERITY_LOW				: color.Yellow		( "LOW"				),
	gl.DEBUG_SEVERITY_NOTIFICATION		: color.Blue 		( "NOTIFICATION"	),
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

	output.Write([]byte(fmt.Sprintf("\n%d: %s of %s severity,\n\traised from %s:\n\t %s\n\n",
							id, types[gltype],
							severitys[severity],
							sources[source],
							 message)))
	if severity != gl.DEBUG_SEVERITY_NOTIFICATION {

		output.Write([]byte("\u001b[30m;1"))
		output.Write([]byte(debug.Stack()))
		output.Write([]byte("\u001b[0m \n\n"))
	}
	/*fmt.Println([]byte(fmt.Sprintf("%d: %s of %s severity, raised from %s: %s\n",
							id, types[gltype],
							severitys[severity],
							sources[source],
							 message))) */
}
// !!!!!!!!!!!!!!!!!!!!!!!!!!!! READ how to set glDebugCallBack(func)!!!!!!!!!!!!!!!!!!!!!!!!
