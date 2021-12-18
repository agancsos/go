package models
import (
	"../../common"
	"encoding/json"
)

type IItem interface {
	ID()             int;
	SetID(a int);
	Name()           string;
	SetName(a string);
	Label()          string;
	SetLabel(a string);
	SetDescription(a string);
	Description()    string;
	SetLastUpdatedDate(a string);
	LastUpdatedDate()string;
	ReloadFromJson(json string);
    ToJsonString()      string;
}

type Item struct {
	id              int    `json:id`;
    name            string `json:name`;
    label           string `json:label`;
    description     string `json:description`;
    lastUpdatedDate string `json:lastUpdatedDate`;
}
func (x *Item) ReloadFromJson(json string) {
	var dict = common.StrToDictionary([]byte(json));
	x.id = dict["id"].(int);
	x.name = dict["name"].(string);
	x.label = dict["label"].(string);
	x.description = dict["description"].(string);
	x.lastUpdatedDate = dict["lastUpdatedDate"].(string);
}

func (x *Item) ToJsonString() string {
	rawJson, _ := json.Marshal(x);
	return common.DictionaryToJsonString(common.StrToDictionary(rawJson));
}

func (x Item) ID()int { return x.id; }
func (x *Item) SetID(a int) {x.id = a; }
func (x Item) Name() string { return x.name; }
func (x *Item) SetName(a string) { x.name = a; }
func (x Item) Label() string { return x.label; }
func (x *Item) SetLabel(a string) { x.label = a; }
func (x *Item) SetDescription(a string) { x.description = a; }
func (x Item) Description()string { return x.description; }
func (x *Item) SetLastUpdatedDate(a string) { x.lastUpdatedDate = a; }
func (x Item) LastUpdatedDate()string { return x.lastUpdatedDate; }

