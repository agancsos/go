package main
import (
	"fmt"
    "os"
    "./classes/common"
    "./classes/sr"
    "path/filepath"
)

func main() {
    var isHelp        = false;
    var SRI           = sr.GetSRInstance();
    var binaryPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
    var ss            = &common.SystemService{};
    ss.ModulePath     = binaryPath;
    SRI.SS            = ss;
    sr.TS = &common.TraceService{ FilePath: fmt.Sprintf("%s/agdo.log", SRI.SS.BuildModuleContainerPath()), TraceLevel: int(common.TL_INFORMATIONAL) };

    for i := 0; i < len(os.Args); i++ {
        switch(os.Args[i]) {
            case "-h", "--help":
                isHelp = true;
                break;
            default:
                break;
        }
    }

    if isHelp {
        sr.HelpMenu();
        os.Exit(0);
    } else {
    }

    os.Exit(0);
}
