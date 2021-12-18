package main
import (
	"os"
	"./classes/common"
	"./classes/sr"
	"./classes/gmon/services"
	"./classes/gmon/components"
	"fmt"
	"strconv"
	"path/filepath"
	"./classes/gmonrpc"
	grpc "google.golang.org/grpc"
	"net"
)

func main() {
	var isHelp		  = false;
	var SRI           = sr.GetSRInstance();
	var binaryPath, _ = filepath.Abs(filepath.Dir(os.Args[0]));
	var ss			  = &common.SystemService{};
	ss.ModulePath     = binaryPath;
	SRI.SS            = ss;
	var cs            = services.GetConfigurationServiceInstance();
	var port, _       = strconv.Atoi(cs.GetProperty("server.servicePort", "3436").(string));
	var traceLevel    = 3;
	traceLevel, _ = strconv.Atoi(cs.GetProperty("traceLevel", "3").(string));
	var logPath = SRI.GetLogFilePath(cs.GetKey("logFilePath", "").(string));

	sr.TS = &common.TraceService{FilePath: logPath, TraceLevel: traceLevel};
	var server		= components.GetMonitorServerInstance();

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
		sr.HelpMenu("SERVER");
		os.Exit(0);
	} else {
		go func() {
			server.Initialize();
		}();

		var opts []grpc.ServerOption;
        listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port));
        if err != nil {
            panic(fmt.Sprintf("Failed to start server.  %s", err));
        }
        var registrar = grpc.NewServer(opts...);
        gmonrpc.RegisterAgentServiceServer(registrar, gmonrpc.NewRpcAgentServiceServer());
		gmonrpc.RegisterMonitorAgentServiceServer(registrar, gmonrpc.NewRpcMonitorAgentServiceServer());
		gmonrpc.RegisterManagementPackServiceServer(registrar, gmonrpc.NewRpcManagementPackServer());
		gmonrpc.RegisterRoleServiceServer(registrar, gmonrpc.NewRpcRoleServiceServer());
        registrar.Serve(listener);
	}

	os.Exit(0);
}
