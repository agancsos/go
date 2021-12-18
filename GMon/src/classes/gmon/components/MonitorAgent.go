package components
import (
	"../services"
	"fmt"
	"../../gmonrpc"
	grpc "google.golang.org/grpc"
	"context"
	"../../sr"
	"strconv"
	"time"
	"../../common"
	"../models"
)

type MonitorAgent struct {
	services	  map[string]services.IService;
	servicePort   int
	serviceName   string
	safeMode	  bool
	endpoint      string
}

var __agent__ *MonitorAgent;

func (x *MonitorAgent) startServices() {
	for key, value := range x.services {
		println(fmt.Sprintf("Starting service (%s)", key));
		sr.TS.TraceInformational(fmt.Sprintf("Starting service (%s)", key), 1);
		go value.Initialize();
	}
}

func (x *MonitorAgent) pulse() {
	var agent = &models.Agent{};
    agent.SetAgentPort(x.servicePort);
	handle, _ := grpc.Dial(x.endpoint, grpc.WithInsecure());
    client := gmonrpc.NewAgentServiceClient(handle);
	client.AddAgent(context.Background(), &gmonrpc.IncomingAgentMessage{Json: common.StrToStrPtr(agent.ToJsonString())}, grpc.WaitForReady(true));
}

func (x *MonitorAgent) rollLogs() {
}

func GetMonitorAgentInstance() *MonitorAgent {
	if __agent__ == nil {
		__agent__ = &MonitorAgent{};
		__agent__.services = map[string]services.IService{};
		__agent__.safeMode = false;
		var cs = services.GetConfigurationServiceInstance();
		__agent__.servicePort = 3435;
		__agent__.serviceName = "agent";
		__agent__.servicePort, _ = strconv.Atoi(cs.GetProperty("servicePort", "3437").(string));
		__agent__.serviceName, _ = cs.GetProperty("serviceName", "").(string);
		var tempValue, _ = strconv.Atoi(cs.GetProperty("safeMode", "0").(string));
		__agent__.safeMode = tempValue > 0;
        __agent__.endpoint = cs.GetProperty("endpoint", "").(string);
	}
	return __agent__;
}

func (x *MonitorAgent) Initialize() {
	x.startServices();
	for ;; {
		x.pulse();
		time.Sleep(30);
	}
}

func (x *MonitorAgent) RegisterService (a string, b services.IService) {}
func (x *MonitorAgent) GetService (a string) services.IService { return nil; }
func (x *MonitorAgent) ContainsService(a string) bool { return true; }
func (x *MonitorAgent) Services() map[string] services.IService { return x.services; }
