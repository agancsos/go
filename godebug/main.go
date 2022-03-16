package main
import (
	"./common"
	"os"
	"./debuggers"
)

var __APPLICATION_NAME__    = "GoDebug";
var __APPLICATION_VERSION__ = "1.0.0";
var __APPLICATION_AUTHOR__  = "Abel Gancsos";
var __APPLICATION_FLAGS__   = map[string]string {
	"-h|--help"          : "Print the help menu",
	"-d|--debugger"      : "Name of the debugger to run",
};

func HelpMenu() {
	println(common.PadLeft("", 80, "#"));
	println(common.PadLeft("# Name     : " + __APPLICATION_NAME__, 79, " ") + "#");
	println(common.PadLeft("# Author   : " + __APPLICATION_AUTHOR__, 79, " ") + "#");
	println(common.PadLeft("# Version  : " + __APPLICATION_VERSION__, 79, " ") + "#");
	println(common.PadLeft("# Flags", 79, " ") + "#");
	for key, value := range __APPLICATION_FLAGS__ {
		println(common.PadLeft("#  " + key + ": " + value, 79, " ") + "#");
	}
	println(common.PadLeft("# Debuggers", 79, " ") + "#");
	for _, value := range debuggers.Debuggers {
		println(common.PadLeft("#  * " + value.GetName() + ": " + value.GetDescription(), 79, " ") + "#");
		println(common.PadLeft("#   * Args", 79, " ") + "#");
		for key, value := range value.GetArguments() {
			println(common.PadLeft("#    * " + key + ": " + value, 79, " ") + "#");
		}
	}
	println(common.PadLeft("", 80, "#"));
}

func main() {
	var isHelp   = false;
	var debugger = "";
	var service  = debuggers.DebuggerService{};
	var args = map[string]string{};

	if len(os.Args) < 2 {
		HelpMenu();
		os.Exit(0);
	}
	for i := 1; i < len(os.Args); i++ {
		switch(os.Args[i]) {
			case "-h", "--help":
				isHelp = true;
				break;
			case "-d", "--debugger":
				debugger = os.Args[i + 1];
				i++;
				break;
			default:
				if i > 1 && i < len(os.Args) - 1 {
					args[os.Args[i]] = os.Args[i + 1];
					i++;
				}
				break;
		}
	}
	if isHelp {
		HelpMenu();
	} else {
		if debugger == "" {
			panic("Debugger name cannot be empty...");
		} else {
			service.Invoke(debugger, args);
		}
	}

	os.Exit(0);
}
