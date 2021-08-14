package main
import (
	"./classes/common"
	"os"
	"path/filepath"
	"strconv"
	"./classes/sr"
	"./classes/helpers"
	"./classes/rpcapi"
	"net"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	var SRI		   = sr.GetSRInstance();
	var binaryPath, _ = filepath.Abs(filepath.Dir(os.Args[0]));
	var ss			= &common.SystemService{};
	ss.ModulePath	 = binaryPath;
	SRI.SS			= ss;
	var cs			= helpers.GetConfigurationServiceInstance();
	var traceLevel	= 3;
	traceLevel,_ = strconv.Atoi(cs.GetProperty("traceLevel", "3").(string));
	var port,_ = strconv.Atoi(cs.GetProperty("serverPort", "4441").(string));
	sr.TS = &common.TraceService{FilePath: SRI.GetLogFilePath(cs.GetProperty("logfilePath", "").(string)), TraceLevel: traceLevel};

	var opts []grpc.ServerOption;
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port));
	if err != nil {
		panic(fmt.Sprintf("Failed to start server.  %s", err));
	}
	var registrar = grpc.NewServer(opts...);
	rpcapi.RegisterRpcApiServer(registrar, &rpcapi.SampleRpcApiServer{});
	registrar.Serve(listener);

	os.Exit(0);
}
