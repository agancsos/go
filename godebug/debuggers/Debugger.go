package debuggers
import (
	"../common"
)

// Constants
var Debuggers = []Debugger {
	&EmptyDebugger{},
	&SystemSearchDebugger{},
	&FileHandleDebugger{},
}
/******************************************************************************/


type Debugger interface {
	GetName()           string
	GetDescription()    string
	GetArguments()      map[string]string
	Invoke(a map[string]string)
}

// DebuggerService
type DebuggerService struct { }
func (x *DebuggerService) Invoke(a string, b map[string]string) {
	var isFound = false;
	for _, cursor := range Debuggers {
		if cursor.GetName() == a {
			isFound = true;
			cursor.Invoke(b);
			break;
		}
	}
	if !isFound {
		panic("Debugger not found!! " + a);
	}
}
/******************************************************************************/

// Dummy debugger
type EmptyDebugger struct {}
func (x *EmptyDebugger) GetName() string { return "EmptyDebugger"; }
func (x *EmptyDebugger) GetDescription() string {
	return "Just a simple debugger that lists a directory.";
}
func (x *EmptyDebugger) Invoke(b map[string]string) {
	println(common.RunCmd("ls -lart " + b["--path"]));
}
func (x *EmptyDebugger) GetArguments() map[string]string {
	return map[string]string {
		"--path" : "Path to list from.  PWD is default.",
	}
}
/******************************************************************************/
