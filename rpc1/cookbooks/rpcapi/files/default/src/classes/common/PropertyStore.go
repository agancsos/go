package common
import (
	"io/ioutil"
	"strings"
)

// PropertyStore interface
type PropertyStore interface {
	GetKeys()        map[string]interface{}
	GetKey(a string) interface{}
	SetProperty(a string, b interface{})
	load(a string)
}
/******************************************************************************/

// PropertyStoreService
type PropertyStoreService struct {}
func (x *PropertyStoreService) LoadFromFile(a string, b PropertyStore) {
	contents, _ := ioutil.ReadFile(a);
	b.load(string(contents));
}
/******************************************************************************/

// EmptyPropertyStore
type EmptyPropertyStore struct{
	properties     map[string]interface{}
}
func (x *EmptyPropertyStore) GetKeys()map[string]interface{} { return x.properties; }
func (x *EmptyPropertyStore) GetKey(a string) interface{} { return x.properties[a]; }
func (x *EmptyPropertyStore) SetProperty(a string, b interface{}) { x.properties[a] = b}
func (x *EmptyPropertyStore) load(a string) {}
/******************************************************************************/

// JsonPropertyStore
type JsonPropertyStore struct{
    properties     map[string]interface{}
}
func (x *JsonPropertyStore) GetKeys()map[string]interface{} { return x.properties; }
func (x *JsonPropertyStore) GetKey(a string) interface{} { return x.properties[a]; }
func (x *JsonPropertyStore) SetProperty(a string, b interface{}) { x.properties[a] = b; }
func (x *JsonPropertyStore) load(a string) {
	var b = StrToDictionary([]byte(a));
	 for key, value := range b {
		x.properties[key] = value;
	}
}
/******************************************************************************/

// PlainPropertyStore
type PlainPropertyStore struct{
    properties     map[string]interface{}
	Delimiter      string
}
func (x *PlainPropertyStore) GetKeys()map[string]interface{} { return x.properties; }
func (x *PlainPropertyStore) GetKey(a string) interface{} { return x.properties[a]; }
func (x *PlainPropertyStore) SetProperty(a string, b interface{}) { x.properties[a] = b; }
func (x *PlainPropertyStore) load(a string) {
	if x.Delimiter == "" { x.Delimiter = ","; }
	for _, line := range strings.Split(a, "\n") {
		comps := strings.Split(line, x.Delimiter);
		x.properties[comps[0]] = comps[1];
	}
}
/******************************************************************************/
