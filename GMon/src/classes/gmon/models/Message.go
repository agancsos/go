package models
import (
    "../../common"
    "encoding/json"
)

type Message struct {
	id              int    `json:id`;
	text            string `json:text`;
    level           int    `json:level`;
    category        int    `json:category`;
    lastUpdatedDate string `json:lastUpdatedDate`;
}
func (x *Message) ReloadFromJson(json string) {
    var dict = common.StrToDictionary([]byte(json));
    x.id = dict["id"].(int);
    x.text = dict["text"].(string);
    x.level = dict["level"].(int);
    x.category = dict["category"].(int);
    x.lastUpdatedDate = dict["lastUpdatedDate"].(string);
}

func (x *Message) ToJsonString() string {
    rawJson, _ := json.Marshal(x);
    return common.DictionaryToJsonString(common.StrToDictionary(rawJson));
}

func (x Message) ID() int { return x.id; }
func (x Message) Text() string { return x.text; }
func (x Message) Level() int { return x.level; }
func (x Message) Category() int { return x.category; }
func (x Message) LastUpdatedDate() string { return x.lastUpdatedDate; }
func (x *Message) SetID(id int) { x.id = id; }
func (x *Message) SetText(text string) { x.text = text; }
func (x *Message) SetLevel(level int) { x.level = level; }
func (x *Message) SetCategory(category int) { x.category = category; }
func (x *Message) SetLastUpdatedDate(ud string) { x.lastUpdatedDate = ud; }
