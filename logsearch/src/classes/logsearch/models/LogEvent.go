package models
import (
	"fmt"
	"../../common"
)

type LogEvent struct {
	id               string
	timestamp        string
	source           string
	data             string
	custom           map[string]string
}

func (x *LogEvent) ReloadFromJson(raw string) {
	var json = common.StrToDictionary([]byte(raw));
	x.id         = fmt.Sprintf("%v", json["id"]);
	x.timestamp  = fmt.Sprintf("%v", json["timestamp"]);
	x.source     = fmt.Sprintf("%v", json["source"]);
	x.data       = fmt.Sprintf("%v", json["data"]);
	x.custom     = json["custom"].(map[string]string);
}

func (x LogEvent) ToJsonString() string {
	var rst  = "{";
	rst      += fmt.Sprintf(`"id": %s`, x.id);
	rst      += fmt.Sprintf(`,"timestamp": "%s"`, x.timestamp);
	rst      += fmt.Sprintf(`, "source": "%s"`, x.source);
	rst      += fmt.Sprintf(`, "data": "%s"`, x.data);
	rst      += fmt.Sprintf(`, "custom": {`);
	var i = 0;
	for k, v := range x.custom {
		if i > 0 { rst += fmt.Sprintf(","); }
		rst += fmt.Sprintf(`"%s": "%s"`, k, v);
	}
	rst      += fmt.Sprintf("}");   
	rst      += "}";
	return rst;
}

func (x LogEvent) ID() string { return x.id; }
func (x LogEvent) Timestamp() string { return x.timestamp; }
func (x LogEvent) Source() string { return x.source; }
func (x LogEvent) Data() string { return x.data; }
func (x LogEvent) CustomFields() map[string]string { return x.custom; }

func (x *LogEvent) SetID(a string) { x.id = a; }
func (x *LogEvent) SetTimestamp(a string) { x.timestamp = a; }
func (x *LogEvent) SetSource(a string) { x.source = a; }
func (x *LogEvent) SetData(a string) { x.data = a; }
func (x *LogEvent) AddCustomField(a string, b string) { x.custom[a] = b; }

