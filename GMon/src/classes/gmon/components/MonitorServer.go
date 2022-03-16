package components
import (
	"../../common"
	"context"
	"../../gmonrpc"
	grpc "google.golang.org/grpc"
	"../services"
	"fmt"
	"strconv"
	"time"
	"../../sr"
	"../models"
)

type MonitorServer struct {
	ds			  *services.DataService
	dbts		  *services.DbTraceService
	services	  map[string]services.IService;
	servicePort   int
	serviceName   string
	disableWC	  bool
	safeMode	  bool
}

var __server__ *MonitorServer;

func (x *MonitorServer) startServices() {
	for key, value := range x.services {
		println(fmt.Sprintf("Starting service (%s)", key));
		sr.TS.TraceInformational(fmt.Sprintf("Starting service (%s)", key), 1);
		go value.Initialize();
	}
}

func (x *MonitorServer) pulse() {
	var packs = services.GetLocalMPServiceInstance().(*services.LocalManagementPackService).GetManagementPacks();
	for _, pack := range packs {
		println(pack.Name());
		var agents = services.GetLocalAgentServiceInstance().(*services.LocalAgentService).GetAgents();
		var agent *models.Agent;
		// Find available Agent
		for _, cursor := range agents { if cursor.State() == 1 { agent = cursor; } }
		if agent != nil {
			go func(a *models.Agent, b *models.ManagementPack) {
				// Send MP to Agent through gRPC call
				var agentProxy = gmonrpc.GetAgentProxy(agent);
				agentProxy.IncomingManagementPack(context.Background(), &gmonrpc.IncomingAgentMessage{Json: common.StrToStrPtr(pack.ToJsonString())}, grpc.WaitForReady(true));
			}(agent, pack);
		}
	}
}

func (x *MonitorServer) rollLogs() {
}

func GetMonitorServerInstance() *MonitorServer {
	if __server__ == nil {
		__server__ = &MonitorServer{};
		__server__.ds = services.GetDataServiceInstance().(*services.DataService);
		__server__.dbts = services.GetDbTraceServiceInstance().(*services.DbTraceService);
		__server__.services = map[string]services.IService{};
		__server__.safeMode = false;
		var cs = services.GetConfigurationServiceInstance();
		__server__.servicePort = 3435;
		__server__.serviceName = "server";
		__server__.disableWC = false;
		__server__.servicePort, _ = strconv.Atoi(cs.GetProperty("servicePort", "3435").(string));
		__server__.serviceName, _ = cs.GetProperty("serviceName", "").(string);
		var tempValue, _ = strconv.Atoi(cs.GetProperty("disableWebSocket", "0").(string));
		__server__.disableWC = tempValue > 0;
		tempValue, _ = strconv.Atoi(cs.GetProperty("safeMode", "0").(string));
		__server__.safeMode = tempValue > 0;
		if !__server__.disableWC {
			__server__.services["RestService"] = services.GetRestServiceInstance();
		}
		if !__server__.disableWC {
		}
	}
	return __server__;
}

func (x *MonitorServer) Initialize() {
	x.startServices();
	for ;; {
		x.pulse();
		time.Sleep(30);
	}
}

func (x *MonitorServer) RegisterService (a string, b services.IService) {}
func (x *MonitorServer) GetService (a string) services.IService { return nil; }
func (x *MonitorServer) ContainsService(a string) bool { return true; }
func (x *MonitorServer) Services() map[string]services.IService { return x.services; }

