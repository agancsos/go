package main
import (
	"./classes/common"
    "os"
	"context"
	"fmt"
    "path/filepath"
    "strconv"
    "./classes/sr"
    "./classes/helpers"
    "./classes/rpcapi"
    "google.golang.org/grpc"
	"strings"
)

func main() {
	var isHelp		= false;
	var SRI		   = sr.GetSRInstance();
	var binaryPath, _ = filepath.Abs(filepath.Dir(os.Args[0]));
	var ss			= &common.SystemService{};
	ss.ModulePath	 = binaryPath;
	SRI.SS			= ss;
	var cs			= helpers.GetConfigurationServiceInstance();
	var traceLevel	= 3;
	var target      = "";
	var method      = "version";
	traceLevel,_ = strconv.Atoi(cs.GetProperty("traceLevel", "3").(string));

	sr.TS = &common.TraceService{FilePath: SRI.GetLogFilePath(cs.GetProperty("logfilePath", "").(string)), TraceLevel: traceLevel};

	if len(os.Args) < 2 {
		sr.HelpMenu();
		os.Exit(0);
	}
	for i := 0; i < len(os.Args); i++ {
		switch(os.Args[i]) {
			case "-h", "--help":
				isHelp = true;
				break;
			case "-t", "--target":
				target = os.Args[i + 1];
				break;
			case "-m", "--method":
				method = os.Args[i + 1];
				break;
			default:
				break;
		}
	}
	if isHelp {
		sr.HelpMenu();
	} else {
		var opts2 grpc.CallOption = grpc.WaitForReady(true);
		handle, err := grpc.Dial(target, grpc.WithInsecure());
		if err != nil {
			panic(fmt.Sprintf("Failed to initialize handle.  %s", err));
		}
		defer handle.Close();
		client := rpcapi.NewRpcApiClient(handle);
		if client == nil {
			panic(fmt.Sprintf("Failed to initialize client."));
		}

		switch (strings.ToLower(method)) {
			case "version":
				var rsp, err = client.GetVersion(context.Background(), &rpcapi.TextMessage{Text:common.StrToStrPtr("")}, opts2);
				if err != nil {
					fmt.Printf("Call failed (%s): %s\n", method, err);
				} else {
					fmt.Printf("Response: %s\n", *rsp.Text);
				}
				break;
			case "hello":
                var rsp, err = client.GetHello(context.Background(), &rpcapi.TextMessage{Text:common.StrToStrPtr("")}, opts2);
                if err != nil {
                    fmt.Printf("Call failed (%s): %s\n", method, err);
                } else {
                    fmt.Printf("Response: %s\n", *rsp.Text);
                }
                break;
			default:
				fmt.Printf("Method not currently implemented (%s)...\n", method);
				break;
		}
	}
	os.Exit(0);
}
