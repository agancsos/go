package main
import (
	"./classes/gmon/services"
	"fmt"
	"./classes/common"
	"./classes/sr"
	"path/filepath"
	"os"
)

func main() {
	var isHelp		  = false;
	var isInstall     = false;
	var binaryPath, _ = filepath.Abs(filepath.Dir(os.Args[0]));
	var component     = "";
	var ss			  = &common.SystemService{};
	ss.ModulePath     = binaryPath;
	//var SRI		      = &sr.SR{SS:ss};
	var ss2		      = services.GetSecurityServiceInstance();
	//var cs            = services.GetConfigurationServiceInstance();
	var installer     = services.GetInstallationServiceInstance();

	for i := 0; i < len(os.Args); i++ {
		switch(os.Args[i]) {
			case "-h", "--help":
				isHelp = true;
				break;
			case "-i", "--install":
				isInstall = true;
				break;
			case "--dba_user":
				installer.SetDbaUsername(os.Args[i + 1]);
				break;
			case "--user":
				installer.SetUsername(os.Args[i + 1]);
				break;
			case "--schema":
				installer.SetSchema(os.Args[i + 1]);
				break;
			case "--server":
				installer.SetServer(os.Args[i + 1]);
				break;
			case "--driver":
				installer.SetDriver(os.Args[i + 1]);
				break;
			case "--dba_pass":
				installer.SetDbaPassword(ss2.GetEncoded(os.Args[i + 1]));
				break;
			case "--pass":
				installer.SetPassword(ss2.GetEncoded(os.Args[i + 1]));
				break;
			case "--debug":
				installer.SetDebug(true);
				break;
			case "--component":
				component = os.Args[i + 1];
				break;
			default:
				break;
		}
	}

	if isHelp {
		sr.HelpMenu("INSTALLER");
		os.Exit(0);
	} else if isInstall {
		// Extract component files
		switch (component) {
			case "server", "agent":
				common.RunCmd(fmt.Sprintf("tar -C /opt/gmon/%s xf ./%s.tar", component));
				break;
			default:
				panic("Component currently not supported in the installer...");
				break;
		}

		// If installing the server component, create the database
		if component == "server" {
			installer.Invoke();
		}
	} else {
	}

	os.Exit(0);
}
