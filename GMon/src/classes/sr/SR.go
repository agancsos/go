package sr
import (
	"../common"
	"net/http"
	b64 "encoding/base64"
	"fmt"
)

// Constants
type TraceCategory int;
const (
    TC_NONE TraceCategory = iota
    TC_SERVICE
)

var TS *common.TraceService;
var APPLICATION_NAME    = "Gancsos Monitor";
var APPLICATION_VERSION = "1.0.0";
var APPLICATION_AUTHOR  = "Gancsos Labs, Inc";
var TOKEN_HEADER_NAME   = "API_TOKEN";
var APPLICATION_FLAGS   = map[string]string {
    "-h|--help"          : "Print the help menu",
    "-i|--install"       : "Install mode",
    "--dba_user"         : "DBA username to be used to create the schema (Install)",
    "--dba_paSS"         : "DBA paSSword to be used to create for the DBA user (Install)",
    "--user"             : "Username to be used for the platform (Install)",
    "--schema"           : "Name of the database for the platform (Install)",
    "--paSS"             : "PaSSword to be used for the user (Install)",
    "--driver"           : "Driver name to be used for the ODBC connection (Install)",
	"--component"        : "Component to install (server|agent) (Install)",
};
var SERVER_FLAGS       = map[string]string {
}
var AGENT_FLAGS        = map[string]string {
}
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
func HelpMenu(a string) {
	var flags = APPLICATION_FLAGS;
	switch (a) {
		case "AGENT":
			flags = AGENT_FLAGS;
			break;
		case "SERVER":
			flags = SERVER_FLAGS;
			break;
		default:
			flags = APPLICATION_FLAGS;
			break;
	}
    println(common.PadLeft("", 80, "#"));
    println(common.PadLeft("# Name     : " + APPLICATION_NAME, 79, " ") + "#");
    println(common.PadLeft("# Author   : " + APPLICATION_AUTHOR, 79, " ") + "#");
    println(common.PadLeft("# Version  : " + APPLICATION_VERSION, 79, " ") + "#");
    println(common.PadLeft("# Flags", 79, " ") + "#");
    for key, value := range flags {
        println(common.PadLeft("#  " + key + ": " + value, 79, " ") + "#");
    }
    println(common.PadLeft("", 80, "#"));
}

func (x *SR) GetConfigurationFile() string {
	return x.SS.BuildModuleContainerPath() + "/resources/config.json";
}

func (x *SR) GetDataFilePath(path string) string {
    if path != "" {
       return path;
	}
    return x.SS.BuildModuleContainerPath() + "/resources/cmserver.db";
}

func (x *SR) GetDataConfigurationFilePath(path string) string {
	if path != "" {
		return path;
	}
    return x.SS.BuildModuleContainerPath() + "/resources/database.json";
}

func (x *SR) GetLogFilePath(path string) string {
	if path != "" {
         return path;
	}
    return x.SS.BuildModuleContainerPath() + "/gmon.log";
}

func (x *SR) GetMediaBasePath(path string) string {
    if path != "" {
         return path;
	}
    return x.SS.BuildModuleContainerPath() + "/media";
}
/*****************************************************************************/

// Helpers
func ExtractApiToken(a *http.Request) string {
	for key, value := range a.Header {
		if key == TOKEN_HEADER_NAME {
			return value[0];
		}
	}
	return "";
}
func RestAuthenticate(a string, b string, c string) string {
	var token = (&common.RestHelper{}).InvokePost(a, map[string]string {"credentials":b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", b, c))),}, map[string]string {});
	return token["token"].(string);
}
/*****************************************************************************/
