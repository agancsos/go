package models
import (
	"fmt"
	"../../common"
	"encoding/json"
	"./discoveries"
	"./datasources"
	"./monitors"
	"./rules"
)

// Variables
type Variable struct {
	name       string
	value      interface{}
}
func (x *Variable) Name() string { return x.name; }
func (x *Variable) SetName(a string) { x.name = a; }
func (x *Variable) Value() interface{} { return x.value; }
func (x *Variable) SetValue(a interface{}) { x.value = a; }

func VariableMapToJson(a map[string]*Variable) string {
	var result = "{";
	var i = 0;
	for _, variable := range a {
		if i > 0 {
			result += ",";
		}
		result += fmt.Sprintf("\"%s\":\"%v\"", variable.Name(), variable.Value());
		i++;
	}
	result += "]";
	return result;
}
/*****************************************************************************/

// Interfaces
type IDiscovery interface {
	ID()               string
	Name()             string
	Type()             string
	DataSource()       datasources.IDataSource
	Invoke()           discoveries.IDiscoveryResponse
	Arguments()        map[string]*Variable
	AddArgument(a *Variable)
	SetDataSource(a datasources.IDataSource)
}

type IMonitor interface {
	ID()              string
	Name()            string
	Type()            string
	Invoke()          monitors.IMonitorResponse
	Arguments()       map[string]*Variable
	DataSource()      datasources.IDataSource
	AddArgument(a *Variable)
	IntervalSeconds() int
	SetDataSource(a datasources.IDataSource)
}

type IRule interface {
	ID()             string
	Name()           string
	Type()           string
	Invoke()         rules.IRuleResponse
	Arguments()      map[string]*Variable
	DataSource()     datasources.IDataSource
	AddArgument(a *Variable)
	SetDataSource(a datasources.IDataSource)
}
/*****************************************************************************/

// Management pack
type ManagementPack struct {
	id              int              `json:id`;
    name            string           `json:name`;
    label           string           `json:label`;
    description     string           `json:description`;
    lastUpdatedDate string           `json:lastUpdatedDate`;
	discoveries     []IDiscovery     `json:discoveries`;
	rules           []IRule          `json:rules`;
	monitors        []IMonitor       `json:monitors`;
}
func (x *ManagementPack) ReloadFromJson(json string) {
    var dict = common.StrToDictionary([]byte(json));
    x.id = dict["id"].(int);
    x.name = dict["name"].(string);
    x.label = dict["label"].(string);
    x.description = dict["description"].(string);
    x.lastUpdatedDate = dict["lastUpdatedDate"].(string);
	x.discoveries = dict["discoveries"].([]IDiscovery);
	x.rules = dict["rules"].([]IRule);
	x.monitors = dict["monitors"].([]IMonitor);
}
func (x *ManagementPack) ToJsonString() string {
    rawJson, _ := json.Marshal(x);
    return common.DictionaryToJsonString(common.StrToDictionary(rawJson));
}
func (x *ManagementPack) AddDiscovery(a IDiscovery) {
	for _, b := range x.discoveries {
		if a.ID() == b.ID() {
			return;
		}
	}
	x.discoveries = append(x.discoveries, a);
}
func (x *ManagementPack) AddMonitor(a IMonitor) {
    for _, b := range x.monitors {
        if a.ID() == b.ID() {
            return;
        }
    }
    x.monitors = append(x.monitors, a);
}
func (x *ManagementPack) AddRule(a IRule) {
    for _, b := range x.rules {
        if a.ID() == b.ID() {
            return;
        }
    }
    x.rules = append(x.rules, a);
}
func (x ManagementPack) ID()int { return x.id; }
func (x *ManagementPack) SetID(a int) {x.id = a; }
func (x ManagementPack) Name() string { return x.name; }
func (x *ManagementPack) SetName(a string) { x.name = a; }
func (x ManagementPack) Label() string { return x.label; }
func (x *ManagementPack) SetLabel(a string) { x.label = a; }
func (x *ManagementPack) SetDescription(a string) { x.description = a; }
func (x ManagementPack) Description()string { return x.description; }
func (x *ManagementPack) SetLastUpdatedDate(a string) { x.lastUpdatedDate = a; }
func (x ManagementPack) LastUpdatedDate()string { return x.lastUpdatedDate; }
func (x ManagementPack) Discoveries() []IDiscovery { return x.discoveries; }
func (x ManagementPack) Rules() []IRule { return x.rules; }
func (x ManagementPack) Monitors() []IMonitor { return x.monitors; }
/*****************************************************************************/
