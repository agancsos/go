package vsphere
import (
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"
)

type VSphereClient struct {
	sessionId             string
	baseEndpoint          string
	username              string
	pat                   string
}

func DictionaryToJsonString (a map[string]interface{}) string {
	var result = "{";
	for key, value := range a {
		result += fmt.Sprintf("\"%s\":\"%v\"", key, value);
	}
	result += "}";
	return result;
}

func NewVSphereClient(path string) (error, *VSphereClient) {
	var rst = &VSphereClient{};
	rst.sessionId    = "";
	rst.baseEndpoint = "";
	_, err    := os.Stat(path);
	if err != nil { return err, nil; }
	raw, err   := ioutil.ReadFile(path);
	var jsonConfig map[string]interface{};
	if err != nil { return err, nil; }
	err = json.Unmarshal(raw, &jsonConfig);
	_, exists := jsonConfig["baseEndpoint"]; if exists { 
		rst.baseEndpoint = jsonConfig["baseEndpoint"].(string);
	}
	_, exists = jsonConfig["username"]; if exists { 
		rst.username = jsonConfig["username"].(string);
	}
	_, exists = jsonConfig["pat"]; if exists { 
		rst.pat = jsonConfig["pat"].(string);
		var decoded, _ = base64.StdEncoding.DecodeString(jsonConfig["pat"].(string));
		var comps      = strings.Split(string(decoded), ":");
		if len(comps) == 2 {
			rst.pat = comps[1];
		}
	}
	return nil, rst;
}

func (x *VSphereClient) authenticate() error {
	if x.sessionId != "" { return nil; }
	var encoded = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", x.username, x.pat)));
	var client       = http.Client{};
	req, err         := http.NewRequest("GET", fmt.Sprintf("%s/api/session", x.baseEndpoint), nil);
	if err != nil { return err; }
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encoded));
	rsp, err := client.Do(req);
	if err != nil { return err; }
	rspData, err := ioutil.ReadAll(rsp.Body);
	if err != nil { return err; }
	x.sessionId = string(rspData);
	return nil;
}

func (x VSphereClient) RawRequest(endpoint string, headers map[string]string) (error, map[string]interface{}) {
	var client       = http.Client{};
	if ! strings.Contains(endpoint, "http") { endpoint = fmt.Sprintf("%s/%s", x.baseEndpoint, endpoint); }
	req, err         := http.NewRequest("GET", endpoint, nil);
	if err != nil { return err, map[string]interface{}{}; }
	for h, v := range headers {
		req.Header.Add(h, v);
	}
	req.Header.Add("vmware-api-session-id", x.sessionId);
	rsp, err := client.Do(req);
	if err != nil { return nil, map[string]interface{}{}; }
	rspData, err := ioutil.ReadAll(rsp.Body);
	if err != nil { return err, map[string]interface{}{}; }
	var jobj map[string]interface{};
	err = json.Unmarshal(rspData, &jobj);
	return err, jobj;
}

func (x VSphereClient) RawPostRequest(method string, endpoint string, headers map[string]string, body map[string]interface{}) (error, map[string]interface{}) {
	var client       = http.Client{};
	if ! strings.Contains(endpoint, "http") { endpoint = fmt.Sprintf("%s/%s", x.baseEndpoint, endpoint); }
	req, err         := http.NewRequest(method, endpoint, bytes.NewBuffer([]byte(DictionaryToJsonString(body))));
	if err != nil { return err, map[string]interface{}{}; }
	for h, v := range headers {
		req.Header.Add(h, v);
	}
	req.Header.Add("vmware-api-session-id", x.sessionId);
	rsp, err := client.Do(req);
	if err != nil { return nil, map[string]interface{}{}; }
	rspData, err := ioutil.ReadAll(rsp.Body);
	if err != nil { return err, map[string]interface{}{}; }
	var jobj map[string]interface{};
	err = json.Unmarshal(rspData, &jobj);
}

func (x VSphereClient) SessionId() string { return x.sessionId; }
func (x VSphereClient) BaseEndpoint() string { return x.baseEndpoint; }

func (x *VSphereClient) SetSessionId(a string) { x.sessionId = a; }
func (x *VSphereClient) SetBaseEndpoint(a string) { x.baseEndpoint = a; }
