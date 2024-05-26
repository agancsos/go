package terratool
var ApplicationName   = "terratool";
var AuthorName        = "Abel Gancsos";
var VersionString     = "1.0.0.0";
var DescriptionString = "Helps generate Terraform templates.";
var Flags             = map[string]string {
    "-h": "Show help menu",
	"--version": "Show version",
    "--dry": "No system changes",
    "--federated": "Generate federated authenticated templates",
    "--provider": "Provider used for the template generation",
    "--sysid": "Application symbol used for tagging",
    "-o": "Full path for the base output",
    "-f": "Full path for input file",
};

