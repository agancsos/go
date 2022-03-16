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
	var binaryPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	var ss			  = &common.SystemService{};
	ss.ModulePath     = binaryPath;
	var SRI		      = &sr.SR{SS:ss};
	var cs            = services.GetConfigurationServiceInstance();
	var traceLevel    = 3;
	var port, _       = strconv.Atoi(cs.GetProperty("agent.servicePort", "3438").(string));
	traceLevel, _ = strconv.Atoi(cs.GetProperty("traceLevel", "3").(string));
	var logPath = SRI.GetLogFilePath(cs.GetKey("logFilePath", "").(string));

	sr.TS = &common.TraceService{FilePath: logPath, TraceLevel: traceLevel};
	var server	  = components.GetMonitorAgentInstance();

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
		sr.HelpMenu("AGENT");
		os.Exit(0);
	} else {
		go func(server *components.MonitorAgent) {
			server.Initialize();
		}(server)

		var opts []grpc.ServerOption;
		listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port));
		if err != nil {
			panic(fmt.Sprintf("Failed to start server.  %s", err));
		}
		var registrar = grpc.NewServer(opts...);
		gmonrpc.RegisterMonitorAgentServiceServer(registrar, gmonrpc.NewRpcMonitorAgentServiceServer());
		registrar.Serve(listener);
	}

	os.Exit(0);
}
