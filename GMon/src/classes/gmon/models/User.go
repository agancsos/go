package models
import (
    "../../common"
    "encoding/json"
)

type User struct {
    id              int    `json:id`;
    firstName       string `json:firstName`;
    lastName        string `json:lastName`;
    createdDate     string `json:createdDate`;
    lastUpdatedDate string `json:lastUpdatedDate`;
	username        string `json:username`;
	password        string `json:password`;
	email           string `json:email`;
	settings        string `json:settings`;
	state           int    `json:state`;
}
func (x *User) ReloadFromJson(json string) {
    var dict = common.StrToDictionary([]byte(json));
    x.id = dict["id"].(int);
    x.firstName = dict["firstName"].(string);
	x.username = dict["username"].(string);
	x.email = dict["email"].(string);
	x.settings = dict["settings"].(string);
	x.password = dict["password"].(string);
	x.state = dict["state"].(int);
    x.lastName = dict["lastName"].(string);
    x.createdDate = dict["createdDate"].(string);
    x.lastUpdatedDate = dict["lastUpdatedDate"].(string);
}

func (x *User) ToJsonString() string {
    rawJson, _ := json.Marshal(x);
    return common.DictionaryToJsonString(common.StrToDictionary(rawJson));
}
func (x User) ID()int { return x.id; }
func (x *User) SetID(a int) { x.id = a; }
func (x User) FirstName() string { return x.firstName; }
func (x *User) SetFirstName(a string) { x.firstName = a; }
func (x User) LastName() string { return x.lastName; }
func (x *User) SetLastName(a string) { x.lastName = a; }
func (x *User) SetCreatedDate(a string) { x.createdDate = a; }
func (x User) CreatedDate()string { return x.createdDate; }
func (x *User) SetLastUpdatedDate(a string) { x.lastUpdatedDate = a; }
func (x User) LastUpdatedDate() string { return x.lastUpdatedDate; }
func (x User) Username() string { return x.username; }
func (x *User) SetUsername(a string) { x.username = a; }
func (x User) Password() string { return x.password; }
func (x *User) SetPassword(a string) { x.password = a; }
func (x User) Email() string { return x.email; }
func (x *User) SetEmail(a string) { x.email = a; }
func (x User) Settings() string { return x.settings; }
func (x *User) SetSettings(a string) { x.settings = a; }
func (x User) State() int { return x.state; }
func (x *User) SetState(a int) { x.state = a; }

