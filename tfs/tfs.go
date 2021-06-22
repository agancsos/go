package main
import (
	"./common"
	"./tfs"
	"os"
)

var __APPLICATION_NAME__    = "TFS";
var __APPLICATION_VERSION__ = "1.0.0";
var __APPLICATION_AUTHOR__  = "Abel Gancsos";
var __APPLICATION_FLAGS__   = map[string]string {
	"-h|--help"          : "Print the help menu",
	"-r|--report"        : "Report to run",
	"-b|--base-endpoint" : "Base endpoint for TFS instance (Required)",
	"-u|--user"          : "Username to use to connect (Required)",
	"-p|--pat"           : "PAT string for the user (Required)",
	"-s|--sprint"        : "Iteration path",
	"-t|--team"          : "Team",
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
	println(common.PadLeft("# Reports", 79, " ") + "#");
	for _, report := range tfs.REPORTS {
        println(common.PadLeft("#  " + report.GetName() + ": " + report.GetDescription(), 79, " ") + "#");
    }
	println(common.PadLeft("", 80, "#"));
}

func main() {
	var isHelp   = false;
	var endpoint = "";
	var username = "";
	var pat      = "";
	var report   = "";
	var team     = "";
	var sprint   = "";

	if len(os.Args) < 2 {
		HelpMenu();
		os.Exit(0);
	}
	for i := 0; i < len(os.Args); i++ {
		switch(os.Args[i]) {
			case "-h", "--help":
				isHelp = true;
				break;
			case "-r", "--report":
				report = os.Args[i + 1];
				break;
			case "-b", "--base-endpoint":
				endpoint = os.Args[i + 1];
				break;
			case "-u", "--user":
				username = os.Args[i + 1];
				break;
			case "-p", "--pat":
				pat = os.Args[i + 1];
				break;
			case "s", "--sprint":
				sprint = os.Args[i + 1];
				break;
			case "-t", "--teamm":
				team = os.Args[i + 1];
				break;
			default:
				break;
		}
	}

	if isHelp {
		HelpMenu();
	}	else {
			if endpoint == "" || username == "" || pat == "" {
				println("Required fields are empty...");
				os.Exit(1);
			}
			var session = tfs.TfsHelper {Endpoint: endpoint, Username: username, Pat: pat, Sprint: sprint, Team: team};
			session.Invoke(report);
	}

	os.Exit(0);
}
