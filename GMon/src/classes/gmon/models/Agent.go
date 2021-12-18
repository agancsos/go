package models
import (
    "../../common"
    "encoding/json"
)

type Agent struct {
    id              int     `json:id`;
    name            string  `json:name`;
    label           string  `json:label`;
    lastUpdatedDate string  `json:lastUpdatedDate`
    publicIP        string  `json:publicIP`
    state           int     `json:state`
	agentPort       int     `json:agentPort`
}
func (x *Agent) ReloadFromJson(json string) {
    var dict = common.StrToDictionary([]byte(json));
    x.id = dict["id"].(int);
    x.name = dict["name"].(string);
    x.label = dict["label"].(string);
    x.publicIP = dict["publicIP"].(string);
    x.state = dict["state"].(int);
	x.agentPort = dict["agentPort"].(int);
    x.lastUpdatedDate = dict["lastUpdatedDate"].(string);
}

func (x Agent) ToJsonString() string {
    rawJson, _ := json.Marshal(x);
    return common.DictionaryToJsonString(common.StrToDictionary(rawJson));
}
func (x Agent) ID()int { return x.id; }
func (x *Agent) SetID(a int) {x.id = a; }
func (x Agent) Name() string { return x.name; }
func (x *Agent) SetName(a string) { x.name = a; }
func (x Agent) Label() string { return x.label; }
func (x *Agent) SetLabel(a string) { x.label = a; }
func (x *Agent) SetLastUpdatedDate(a string) { x.lastUpdatedDate = a; }
func (x Agent) LastUpdatedDate() string { return x.lastUpdatedDate; }
func (x Agent) PublicIP() string { return x.publicIP; }
func (x *Agent) SetPublicIP(a string) { x.publicIP = a; }
func (x Agent) State() int { return x.state; }
func (x *Agent) SetState(a int) { x.state = a; }
func (x *Agent) SetAgentPort(a int) { x.agentPort = a; }
func (x Agent) AgentPort() int { return x.agentPort; }
