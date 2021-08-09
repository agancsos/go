package sr
import (
	"../common"
	"fmt"
)

// Constants and shared global variables
type TraceCategory int;
const (
    TC_NONE TraceCategory = iota
    TC_SERVICE
)

var TS *common.TraceService;
var APPLICATION_NAME    = "Abel Gancsos DigitalOcean Project";
var APPLICATION_VERSION = "1.0.0.0";
var APPLICATION_AUTHOR  = "Abel Gancsos";
var APPLICATION_FLAGS   = map[string]string {
    "-h|--help"          : "Print the help menu",
	"-v|--version"       : "Print the version",
	"--op"               : "Operation to perform",
	"-p|--package"       : "Package name",
	"-r|--registry"      : "URL to the package registry",
};
var SUPPORTED_OPERATIONS = map[string]string {
	"get"             : "List packages",
	"install"         : "Install a package",
	"purge"           : "Remove a package",
	"update"          : "Update local cache",
	"upgrade"         : "Upgrade packages",
	"upload"          : "Upload a package to the registry",
};
/*****************************************************************************/

// Static resources
type SR struct {
	SS    *common.SystemService
}
var __sr__ *SR;
func GetSRInstance() *SR {
	if __sr__ == nil {
		__sr__ = &SR{};
	}
	return __sr__;
}

func HelpMenu() {
    fmt.Printf("%s#\n", common.PadLeft("", 80, "#"));
    fmt.Printf("%s#\n", common.PadLeft("# Name     : " + APPLICATION_NAME, 80, " "));
    fmt.Printf("%s#\n", common.PadLeft("# Author   : " + APPLICATION_AUTHOR, 80, " "));
    fmt.Printf("%s#\n", common.PadLeft("# Version  : " + APPLICATION_VERSION, 80, " "));
    fmt.Printf("%s#\n", common.PadLeft("# Flags", 80, " "));
    for key, value := range APPLICATION_FLAGS {
        fmt.Printf("%s#\n", common.PadLeft("#  " + key + ": " + value, 80, " "));
    }
	fmt.Printf("%s#\n", common.PadLeft("# Operations", 80, " "));
    for key, value := range SUPPORTED_OPERATIONS {
        fmt.Printf("%s#\n", common.PadLeft("#  " + key + ": " + value, 80, " "));
    }
    fmt.Printf("%s#\n", common.PadLeft("", 80,"#"));
}

func (x *SR) GetConfigurationFile() string {
	return x.SS.BuildModuleContainerPath() + "/resources/config.json";
}
/*****************************************************************************/

