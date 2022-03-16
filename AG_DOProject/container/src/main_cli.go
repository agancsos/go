package main
import (
	"fmt"
	"os"
	"./classes/common"
	"./classes/sr"
	"./classes/helpers"
	"path/filepath"
	"strconv"
)

func main() {
	var isHelp			  = false;
	var isVersion         = false;
	var SRI			      = sr.GetSRInstance();
	var ss				  = &common.SystemService{};
	var cs				  = helpers.GetConfigurationServiceInstance();
	var binaryPath, _     = filepath.Abs(filepath.Dir(os.Args[0]));
	ss.ModulePath		  = binaryPath;
	SRI.SS				  = ss;
	var traceLevel, err   = strconv.Atoi(cs.GetProperty("traceLevel", "3").(string));
	var operation		  = "";
	var packageName		  = "";
	var repoName          = "";
	var repoUrl           = "";
	var cache             = helpers.GetRegistryCacheInstance();

	if err != nil { fmt.Printf("%s\n", err); }
	sr.TS = &common.TraceService{ FilePath: fmt.Sprintf("%s/agdo.log", SRI.SS.BuildModuleContainerPath()), TraceLevel: traceLevel };

	for i := 0; i < len(os.Args); i++ {
		switch(os.Args[i]) {
			case "--op":
				operation = os.Args[i + 1];
				break;
			case "-h", "--help":
				isHelp = true;
				break;
			case "-v", "--version":
				isVersion = true;
				break;
			case "-p", "--package":
				packageName = os.Args[i + 1];
				break;
			case "-r", "--registry":
				cache.SetRegistryUrl(os.Args[i + 1]);
				break;
			default:
				break;
		}
	}

	if isHelp {
		sr.HelpMenu();
		os.Exit(0);
	} else if isVersion {
		fmt.Printf("%s v. %s\n", sr.APPLICATION_NAME, sr.APPLICATION_VERSION);
	} else {
		cache.Invoke(operation, packageName, repoName, repoUrl);
	}

	os.Exit(0);
}
