package main
import (
	"os"
	"./data"
	"./common"
	"errors"
)
var __VERSION_STRING__      = "1.0.0"
var __APPLICATION_NAME__    = "dataexec"
var __AUTHOR_NAME__         = "Abel Gancsos"
var __APPLICATION_OPTIONS__ = map[string]string {
	"-h|--help"              : "Prints the help menu",
	"-c|--connection-string" : "Connection string to database",
	"-u|--username"          : "Username to use to connect",
	"-p|--password"          : "Password for the database user",
	"-o|--op"                : "Operation to perform (query|execute)",
	"-s|--sql"               : "SQL statement to run",
}

func HelpMenu() {
	println(common.PadLeft("", 80, "#"));
	println(common.PadLeft("# Name    : " + __APPLICATION_NAME__, 79, " ") + "#");
	println(common.PadLeft("# Author  : " + __AUTHOR_NAME__, 79, " ") + "#");
	println(common.PadLeft("# Version : v. " + __VERSION_STRING__, 79, " ") + "#");
	println(common.PadLeft("# Flags   ", 79, " ") + "#");
	for key, value := range __APPLICATION_OPTIONS__ {
		println(common.PadLeft("#  " + key + " : " + value, 79, " ") + "#");
	}
	println(common.PadLeft("# Supported Providers", 79, " ") + "#");
	for key, value := range data.SupportedProviders {
		println(common.PadLeft("#  " + key + " : " + value, 79, " ") + "#");
	}
	println(common.PadLeft("", 80, "#"));
}

func main() {
	_ = data.DataTable{};
	var connectionString = "";
	var username         = "";
	var password         = "";
	var operation        = "";
	var sql              = "";
	var isHelp           = false;

	if len(os.Args) < 2 {
		HelpMenu();
	} else {
		for i := 0; i < len(os.Args); i++ {
			switch os.Args[i] {
				case "-h", "--help":
					isHelp = true;
					break;
				case "-c", "--connection-string":
					connectionString = os.Args[i + 1];
					break;
				case "-u", "--username":
					username = os.Args[i + 1];
					break;
				case "-p", "--password":
					password = os.Args[i + 1];
					break;
				case "-o", "--op":
					operation = os.Args[i + 1];
					break;
				case "-s", "--query":
					sql = os.Args[i + 1];
					break;
				default :
					break;
			}
		}
		if (isHelp) {
			HelpMenu();
		} else {
			if connectionString == "" { panic(errors.New("Connection string cannot be empty...")); }
			if operation == "" { panic(errors.New("Operation cannot be empty...")); }
			if sql == "" { panic(errors.New("Query cannot be empty...")); }
			conn := data.CreateConnection(connectionString, username, password);
			switch operation {
				case "query":
					var result = conn.Query(sql);
					for _ = range result.GetRows() {
					}
					break;
				case "execute":
					var result = conn.RunQuery(sql);
					println(result);
					break;
				default:
					panic(errors.New("Invalid operation specified: " + operation));
			}
		}
	}
	os.Exit(0);
}
