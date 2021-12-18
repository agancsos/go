package models
import (
    "../../common"
    "encoding/json"
)

type Person struct {
    id              int    `json:id`;
    firstName       string `json:firstName`;
    lastName        string `json:lastName`;
    createdDate     string `json:createdDate`;
    lastUpdatedDate string `json:lastUpdatedDate`;
}
func (x *Person) ReloadFromJson(json string) {
    var dict = common.StrToDictionary([]byte(json));
    x.id = dict["id"].(int);
    x.firstName = dict["firstName"].(string);
    x.lastName = dict["lastName"].(string);
    x.createdDate = dict["createdDate"].(string);
    x.lastUpdatedDate = dict["lastUpdatedDate"].(string);
}

func (x *Person) ToJsonString() string {
    rawJson, _ := json.Marshal(x);
    return common.DictionaryToJsonString(common.StrToDictionary(rawJson));
}
func (x Person) ID()int { return x.id; }
func (x Person) SetID(a int) {x.id = a; }
func (x Person) FirstName() string { return x.firstName; }
func (x *Person) SetFirstName(a string) { x.firstName = a; }
func (x Person) LastName() string { return x.lastName; }
func (x *Person) SetLastName(a string) { x.lastName = a; }
func (x *Person) SetCreatedDate(a string) { x.createdDate = a; }
func (x Person) CreatedDate()string { return x.createdDate; }
func (x Person) SetLastUpdatedDate(a string) { x.lastUpdatedDate = a; }
func (x Person) LastUpdatedDate() string { return x.lastUpdatedDate; }

