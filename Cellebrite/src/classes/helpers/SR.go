package helpers
import (
	"strings"
	"unicode"
)

// Constants
var APPLICATION_NAME    = "Cellebrite Interview Practice";
var APPLICATION_VERSION = "1.0.0";
var APPLICATION_AUTHOR  = "Abel Gancsos";
var APPLICATION_FLAGS   = map[string]string {
    "-h|--help"          : "Print the help menu",
	"-t|--tool"          : "Tool to run.",
	"-o|--op"            : "Operation of the tool to run.",
	"-p|--path"          : "Full path of resource for tool operation.",
	"--port"             : "Port for the REST API.  Default is 3434.",
};
/*****************************************************************************/

// Static resources
func HelpMenu() {
    println(PadLeft("", 80, "#"));
    println(PadLeft("# Name     : " + APPLICATION_NAME, 79, " ") + "#");
    println(PadLeft("# Author   : " + APPLICATION_AUTHOR, 79, " ") + "#");
    println(PadLeft("# Version  : " + APPLICATION_VERSION, 79, " ") + "#");
    println(PadLeft("# Flags", 79, " ") + "#");
    for key, value := range APPLICATION_FLAGS {
        println(PadLeft("#  " + key + ": " + value, 79, " ") + "#");
    }
    println(PadLeft("", 80, "#"));
}
/*****************************************************************************/

// Helpers
func PadRight(str string, le int, pad string) string {
	if len(str) > le {
		return str[0:le];
	}
	result := "";
	for i := len(str); i < le; i++ {
		result += pad;
	}
	return result + str;
}

func CleanString(a string) string {
	var result = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, a);
	return result;
}

func PadLeft(str string, le int, pad string) string {
	if len(str) > le {
		return str[0:le];
	}
	result := "";
	for i := len(str); i < le; i++ {
		result += pad;
	}
	return str + result;
}
func ArgsToDictionary(a []string) map[string]string {
	var result = map[string]string{};
	for i := 0; i < len(a) - 1; i++ {
		result[a[i]] = a[i + 1];
	}
	return result;
}
/*****************************************************************************/
