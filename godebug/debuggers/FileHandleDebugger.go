package debuggers
import (
    "fmt"
	"strings"
	"../common"
)

type FileHandleDebugger struct {}
func (x *FileHandleDebugger) GetName() string { return "FileHandle"; }
func (x *FileHandleDebugger) GetDescription() string {
    return "List file handles via processes";
}
func (x *FileHandleDebugger) Invoke(b map[string]string) {
	var search = b["--search"];
	var rsp = strings.Split(common.RunCmd("ps -efl"), "\n");
	for _, line := range rsp {
		if strings.Contains(line, search) || search == "" {
			println(fmt.Sprintf("* %v", line));
		}
	}
}
func (x *FileHandleDebugger) GetArguments() map[string]string {
    return map[string]string {
		"--search"  : "Keyword to search for",
    }
}
