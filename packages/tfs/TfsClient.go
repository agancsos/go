package tfs
import (
	"bytes"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

type TfsClient struct {
	BaseEndpoint                 string   `json:baseEndpoint`
	Username                     string   `json:username`
	PAT                          string   `json:pat`
	Team                         string   `json:team`
}

func DictionaryToJsonString (a map[string]string) string {
	var result = "{";
	for key, value := range a {
		result += fmt.Sprintf("\"%s\":\"%v\"", key, value);
	}
	result += "}";
	return result;
}

func NewClient(path string) (*TfsClient, error) {
	_, err    := os.Stat(path);
	if err != nil { return nil, err; }
	raw, err   := ioutil.ReadFile(path);
	if err != nil { return nil, err; }
	var client   *TfsClient;
	err = json.Unmarshal(raw, &client);
	return client, err;
}


func (x TfsClient) RawRequest(path string) (string, error) {
	var client       = http.Client{};
	req, err         := http.NewRequest("GET", path, nil);
	if err != nil { return "", err; }
	req.SetBasicAuth(x.Username, x.PAT);
	rsp, err := client.Do(req);
	if err != nil { return "", err; }
	rspData, err := ioutil.ReadAll(rsp.Body);
	return string(rspData), err;
}

func (x TfsClient) TfsRequest(path string) (map[string]interface{}, error) {
	var rsp, err  = x.RawRequest(fmt.Sprintf("%s", path));
	if err != nil { return nil, err; }
	var jobj map[string]interface{};
	err = json.Unmarshal([]byte(rsp), &jobj);
	return jobj, err;
}

func (x TfsClient) TfsPostRequest(path string, body map[string]string) (map[string]interface{}, error) {
	var client      = http.Client{};
	req, err        := http.NewRequest("POST", path, bytes.NewBuffer([]byte(DictionaryToJsonString(body))));
	req.Header.Add("Content-Type", "application/json");
	req.Header.Add("Accepts", "application/json");
	req.SetBasicAuth(x.Username, x.PAT);
	rsp, err        := client.Do(req);
	if err != nil { return nil, err; }
	rspData, err    := ioutil.ReadAll(rsp.Body);
	if err != nil { return nil, err; }
	var jobj map[string]interface{};
	err            = json.Unmarshal(rspData, &jobj);
	return jobj, err;
}

