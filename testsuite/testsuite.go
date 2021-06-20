package main
import (
	"./common"
	"./testsuite"
	"os"
)
var __APPLICATION_NAME__    = "TestSuite";
var __APPLICATION_VERSION__ = "1.0.0";
var __APPLICATION_AUTHOR__  = "Abel Gancsos";
var __APPLICATION_FLAGS__   = map[string]string {
	"-h|--help" : "Print the help menu",
	"-o|--op"   : "Operation to perform (list or run)",
	"-t|--test" : "Name of the Unit Test to run",
};

func HelpMenu() {
	println(common.PadLeft("", 80, "#"));
	println(common.PadLeft("# Name     : " + __APPLICATION_NAME__, 79, " ") + "#");
	println(common.PadLeft("# Author   : " + __APPLICATION_AUTHOR__, 79, " ") + "#");
	println(common.PadLeft("# Version  : " + __APPLICATION_VERSION__, 79, " ") + "#");
	println(common.PadLeft("# Flags", 79, " ") + "#");
	for key, value := range __APPLICATION_FLAGS__ {
		println(common.PadLeft("#  " + key + "   : " + value, 79, " ") + "#");
	}
	println(common.PadLeft("", 80, "#"));
}

func main() {
	var isHelp    = false;
	var operation = "";
	var testName  = "";
	var ts testsuite.TestSuite;
	ts.Initialize();

	if len(os.Args) < 2{
		HelpMenu();
		os.Exit(0);
	} else {
		for i := 0; i < len(os.Args); i++ {
			switch(os.Args[i]) {
				case "-h", "--help":
					isHelp = true;
					break;
				case "-o", "--op":
					operation = os.Args[i + 1];
					break;
				case "-t", "--test":
					testName = os.Args[i + 1];
					break;
				default:
					break;
			}
		}
	}

	if isHelp {
		HelpMenu();
	} else {
		switch(operation) {
			case "list":
				ts.ListTests();
				break;
			case "run":
				if testName == "" {
					panic("Test name cannot be empty...");
				}
				if unitTest, ok := testsuite.UnitTests[testName]; ok {
					println(ts.Invoke(unitTest));
				}
				break;
			default:
				panic("Operation not supported at this time...");
				break;
		}
	}
	os.Exit(0);
}

