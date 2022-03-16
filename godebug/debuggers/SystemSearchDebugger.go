package debuggers
import (
	"strings"
	"fmt"
	"os"
	"io/ioutil"
)

// Not really a debugger, mostly just an example of IO
type SystemSearchDebugger struct {}
func (x *SystemSearchDebugger) GetName() string { return "SystemSearch"; }
func (x *SystemSearchDebugger) GetDescription() string {
    return "Search for text in file name and contents";
}
func (x *SystemSearchDebugger) searchPath(path string, search string) {
	var stat, _ = os.Stat(path);
	if stat.Mode().IsRegular() {
		return;
	}
	var files, _ = os.ReadDir(path);
	for _, file := range files {
		// Check file name
		if strings.Contains(file.Name(), search) {
			println(fmt.Sprintf("* %s", file));
			x.searchPath(fmt.Sprintf("%s/%s", path, file.Name()), search);
			continue;
		}
		// CHeck file contents
		var contents, _ = ioutil.ReadFile(fmt.Sprintf("%s/%s", path, file.Name()));
		if strings.Contains(string(contents), search) {
			println(fmt.Sprintf("* %s/%s", path, file.Name()));
		}
		x.searchPath(fmt.Sprintf("%s/%s", path, file.Name()), search);
	}
}
func (x *SystemSearchDebugger) Invoke(b map[string]string) {
	var search = b["--search"];
	var basePath = b["--base-path"];
	if basePath == "" { basePath = "/"; }
	x.searchPath(basePath, search);
}
func (x *SystemSearchDebugger) GetArguments() map[string]string {
    return map[string]string {
        "--search"    : "Text to search for.",
		"--base-path" : "Path to start search.",
    }
}
