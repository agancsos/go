package models
import (
	"encoding/json"
	"../../common"
)

type Role struct {
    id      int      `json:id`
    name    string   `json:name`
}

func (x *Role) ReloadFromJson (json string) {
	var dict = common.StrToDictionary([]byte(json));
    x.id = dict["id"].(int);
    x.name = dict["name"].(string);
}

func (x *Role) ToJsonString() string {
    rawJson, _ := json.Marshal(x);
    return common.DictionaryToJsonString(common.StrToDictionary(rawJson));
}
func (x Role) ID() int { return x.id; }
func (x Role) Name() string { return x.name; }
func (x *Role) SetID(a int) { x.id = a; }
func (x *Role) SetName(a string) { x.name = a; }
