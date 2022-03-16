package sr
import (
	"../common"
)

// Constants
type TraceCategory int;
const (
    TC_NONE TraceCategory = iota
    TC_SERVICE
)

var TS *common.TraceService;
var APPLICATION_NAME    = "Abel Gancsos DigitalOcean Test";
var APPLICATION_VERSION = "1.0.0.0";
var APPLICATION_AUTHOR  = "Abel Gancsos";
var APPLICATION_FLAGS   = map[string]string {
    "-h|--help"          : "Print the help menu",
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
    println(common.PadLeft("", 80, "#"));
    println(common.PadLeft("# Name     : " + APPLICATION_NAME, 79, " ") + "#");
    println(common.PadLeft("# Author   : " + APPLICATION_AUTHOR, 79, " ") + "#");
    println(common.PadLeft("# Version  : " + APPLICATION_VERSION, 79, " ") + "#");
    println(common.PadLeft("# Flags", 79, " ") + "#");
    for key, value := range APPLICATION_FLAGS {
        println(common.PadLeft("#  " + key + ": " + value, 79, " ") + "#");
    }
    println(common.PadLeft("", 80, "#"));
}

func (x *SR) GetConfigurationFile() string {
	return x.SS.BuildModuleContainerPath() + "/resources/config.json";
}
/*****************************************************************************/

// Helpers
/*****************************************************************************/
