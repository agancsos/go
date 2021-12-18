package services
import (
	"../models"
	"../../sr"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"../../common"
)

// Interface
type IAgentService interface {
	GetAgent(id int)					 *models.Agent
	GetAgents()						     []*models.Agent
	AddAgent(a *models.Agent)	         bool
	UpdateAgent(a *models.Agent)         bool
	RemoveAgent(a *models.Agent)         bool
	Contains(hostname string)		     bool
}
/*****************************************************************************/

// Local service
type LocalAgentService struct {
	pollingInterval  time.Duration
	ds				 *DataService
}
var __local_node_service__ *LocalAgentService;
func GetLocalAgentServiceInstance() IService {
	if __local_node_service__ == nil {
		__local_node_service__ = &LocalAgentService{};
		__local_node_service__.ds = GetDataServiceInstance().(*DataService);
		__local_node_service__.pollingInterval = 30;
		var servicesAgent = (GetConfigurationServiceInstance()).GetProperty("services", "");
		if servicesAgent != nil {
			var nsAgent = servicesAgent.(map[string]interface{});
			if nsAgent != nil {
				var interval = nsAgent["interval"].(string);
				intervalValue, _ := time.ParseDuration(interval);
				__local_node_service__.pollingInterval = intervalValue;
			}
		}
	}
	return __local_node_service__;
}

func (x *LocalAgentService) AcceptIncomingMP(mp *models.ManagementPack) {
    // Process Discoveries
    for _, d := range mp.Discoveries() {
        go func(d models.IDiscovery) {
            var rsp = d.Invoke();
			// Add result to discovery runs
			x.ds.RunServiceQuery(fmt.Sprintf("INSERT INTO DISCOVERY_RUNS (DISCOVERY_ID, INPUT_MANIFEST, RESPONSE_MANIFEST) VALUES ('%d', '%s', '%s')",
				d.ID(), models.VariableMapToJson(d.Arguments()), fmt.Sprintf("{\"response\":\"%v\"}", rsp)));
        }(d);
    }

    // Process Rules
    for _, r := range mp.Rules() {
        go func(r models.IRule) {
            var rsp = r.Invoke();
			// Add result to rule runs
			x.ds.RunServiceQuery(fmt.Sprintf("INSERT INTO RULE_RUNS (RULE_ID, INPUT_MANIFEST, RESPONSE_MANIFEST) VALUES ('%d', '%s', '%s')",
				r.ID(), models.VariableMapToJson(r.Arguments()), fmt.Sprintf("{\"response\":\"%v\"}", rsp)));
        }(r);
    }

    // Process Monitors
    for _, m := range mp.Monitors() {
        go func(m models.IMonitor) {
            for ;; {
                var rsp = m.Invoke();
				// Add result to monitor runs
				x.ds.RunServiceQuery(fmt.Sprintf("INSERT INTO MONITOR_RUNS (MONITOR_ID, INPUT_MANIFEST, RESPONSE_MANIFEST) VALUES ('%d', '%s', '%s')",
					m.ID(), models.VariableMapToJson(m.Arguments()), fmt.Sprintf("{\"response\":\"%v\"}", rsp)));
                if m.IntervalSeconds() < 1 { break; }
                time.Sleep(time.Duration(m.IntervalSeconds()));
            }
        }(m);
    }
}

func (x *LocalAgentService) GetAgent(id int) *models.Agent {
	var result *models.Agent;
	sr.TS.TraceVerbose(fmt.Sprintf("AgentService: ting node (%d)", id), int(sr.TC_SERVICE));
	var rows = x.ds.ServiceQuery(fmt.Sprintf("SELECT * FROM AGENTS WHERE AGENT_ID = '%d'", id));
	if len(rows.Rows()) == 1 {
		var row = rows.Rows()[0];
		result = &models.Agent{};
		id, _ = strconv.Atoi(row.Column("AGENT_ID").Value());
		result.SetID(id);
		result.SetName(row.Column("AGENT_HOSTNAME").Value());
		result.SetPublicIP(row.Column("AGENT_PUBLICIP").Value());
		state, _ := strconv.Atoi(row.Column("AGENT_STATE").Value());
		port, _ := strconv.Atoi(row.Column("AGENT_PORT").Value());
		result.SetAgentPort(port);
		result.SetState(state);
	}
	sr.TS.TraceVerbose(fmt.Sprintf("AgentService: Done fetching node (%d)", id), int(sr.TC_SERVICE));
	return result;
}

func (x *LocalAgentService) GetAgents() []*models.Agent {
	var results = []*models.Agent{};
	sr.TS.TraceVerbose("AgentService: ting nodes", int(sr.TC_SERVICE));
	var rows = x.ds.ServiceQuery("SELECT * FROM AGENTS").Rows();
	for _, row := range rows {
		id, _ := strconv.Atoi(row.Column("AGENT_ID").Value());
		results = append(results, x.GetAgent(id));
	}
	sr.TS.TraceVerbose("AgentService: Done fetching nodes", int(sr.TC_SERVICE));
	return results;
}

func (x *LocalAgentService) AddAgent(a *models.Agent) bool {
	sr.TS.TraceVerbose(fmt.Sprintf("AgentService: Adding node (%s)", a.Name()), int(sr.TC_SERVICE));
	var result = false;
	if !x.Contains(a.PublicIP()) {
		result = x.ds.RunServiceQuery(fmt.Sprintf("INSERT INTO AGENTS (AGENT_HOSTNAME, AGENT_PUBLICIP, AGENT_STATE, AGENT_PORT) VALUES ('%s', '%s', '%d', '%d')", a.Name(), a.PublicIP(), a.State(), a.AgentPort()));
	} else {
		result = x.UpdateAgent(a);
	}
	sr.TS.TraceVerbose(fmt.Sprintf("AgentService: Done updating node (%s)", a.Name()), int(sr.TC_SERVICE));
	return result;
}

func (x *LocalAgentService) UpdateAgent(a *models.Agent) bool {
	sr.TS.TraceVerbose(fmt.Sprintf("AgentService: Updating node (%d)", a.ID()), int(sr.TC_SERVICE));
	var result = false;
	result = x.ds.RunServiceQuery(fmt.Sprintf("UPDATE AGENTS SET AGENT_HOSTNAME = '%s', AGENT_PUBLICIP = '%s', AGENT_PORT = '%d' LAST_UPDATED_DATE = CURRENT_TIMESTAMP, AGENT_STATE = '%d' WHERE AGENT_ID = '%d'",
		a.Name(), a.PublicIP(), a.AgentPort(), a.State()));
	sr.TS.TraceVerbose(fmt.Sprintf("AgentService: Done updating node (%d)", a.ID()), int(sr.TC_SERVICE));
	return result;
}

func (x *LocalAgentService) RemoveAgent(a *models.Agent) bool {
	sr.TS.TraceVerbose(fmt.Sprintf("AgentService: Removing node (%d)", a.ID()), int(sr.TC_SERVICE));
	return x.ds.RunServiceQuery(fmt.Sprintf("DELETE FROM AGENTS WHERE AGENT_ID = '%d'", a.ID()));
}

func (x *LocalAgentService) Contains(hostname string) bool {
	sr.TS.TraceVerbose(fmt.Sprintf("AgentService: Checking if node exists (%s)", hostname), int(sr.TC_SERVICE));
	return len(x.ds.ServiceQuery(fmt.Sprintf("SELECT 1 FROM AGENTS WHERE AGENT_HOSTNAME = '%s' OR AGENT_PUBLICIP = '%s'", hostname, hostname)).Rows()) == 1;
}

func (x *LocalAgentService) poll() {
}

func (x *LocalAgentService) Initialize() {
	for ;; {
		x.poll();
		time.Sleep(x.pollingInterval);
	}
}
/*****************************************************************************/

// Rest service
type RestAgentService struct {
	lns     *LocalAgentService
	auth    *LocalAuthenticationService
}
var __rest_node_service__ *RestAgentService;
func GetRestAgentServiceInstance() *RestAgentService {
	if __rest_node_service__ == nil {
		__rest_node_service__ = &RestAgentService{};
		__rest_node_service__.lns = GetLocalAgentServiceInstance().(*LocalAgentService);
		__rest_node_service__.auth = GetLocalAuthService();
	}
	return __rest_node_service__;
}

func (x *RestAgentService) GetAgent(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "POST") {
		return;
	}
    data, _ := ioutil.ReadAll(r.Body);
	id, _ := strconv.Atoi(string(data));
	var node = x.lns.GetAgent(id);
    w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", x.lns.AddAgent(node))));
}

func (x *RestAgentService) GetAgents(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "GET") {
		return;
	}
	var rsp = "{\"nodes\":[";
	var nodes = x.lns.GetAgents();
	for i, node := range nodes {
		if i > 0 {
			rsp += ",";
		}
		rsp += node.ToJsonString();
	}
	rsp += "]}";
    w.Write([]byte(rsp));
}

func (x *RestAgentService) AddAgent(w http.ResponseWriter, r *http.Request) {
	if !common.EnsureRestMethod(r, "POST") {
		return;
	}
	var node *models.Agent;
	data, _ := ioutil.ReadAll(r.Body);
	json.Unmarshal(data, &node);
	w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", common.BoolToInt(x.lns.AddAgent(node)))));
}

func (x *RestAgentService) UpdateAgent(w http.ResponseWriter, r *http.Request) {
	if !common.EnsureRestMethod(r, "POST") {
		return;
	}
	var node *models.Agent;
    data, _ := ioutil.ReadAll(r.Body);
    json.Unmarshal(data, &node);
    w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", common.BoolToInt(x.lns.UpdateAgent(node)))));
}

func (x *RestAgentService) RemoveAgent(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "POST") {
		return;
	}
	var node *models.Agent;
    data, _ := ioutil.ReadAll(r.Body);
    json.Unmarshal(data, &node);
    w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", common.BoolToInt(x.lns.RemoveAgent(node)))));
}

func (x *RestAgentService) Contains(w http.ResponseWriter, r *http.Request) {
	if !EnsureAuthenticated(w, r, "POST") {
		return;
	}
	var node map[string]interface{};
    data, _ := ioutil.ReadAll(r.Body);
    json.Unmarshal(data, &node);
    w.Write([]byte(fmt.Sprintf("{\"result\":\"%d\"}", common.BoolToInt(x.lns.Contains(node["hostname"].(string))))));
}

func (x *RestAgentService) Initialize() {
	http.HandleFunc("/agent/get/", x.GetAgent);
	http.HandleFunc("/agent/list/", x.GetAgents);
	http.HandleFunc("/agent/add/", x.AddAgent);
	http.HandleFunc("/agent/update/", x.UpdateAgent);
	http.HandleFunc("/agent/delete/", x.RemoveAgent);
}
/*****************************************************************************/

