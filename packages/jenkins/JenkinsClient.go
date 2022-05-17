package jenkins
import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

type JenkinsClient struct {
	BaseEndpoint                 string   `json:baseEndpoint`
	Username                     string   `json:username`
	PAT                          string   `json:pat`
}

func NewClient(path string) (*JenkinsClient, error) {
	_, err    := os.Stat(path);
	if err != nil { return nil, err; }
	raw, err   := ioutil.ReadFile(path);
	if err != nil { return nil, err; }
	var client   *JenkinsClient;
	err = json.Unmarshal(raw, &client);
	return client, err;
}


func (x JenkinsClient) RawRequest(path string) (string, error) {
	var client       = http.Client{};
	req, err         := http.NewRequest("GET", path, nil);
	if err != nil { return "", err; }
	req.SetBasicAuth(x.Username, x.PAT);
	rsp, err := client.Do(req);
	if err != nil { return "", err; }
	rspData, err := ioutil.ReadAll(rsp.Body);
	return string(rspData), err;
}

func (x JenkinsClient) JenkinsRequest(path string) (map[string]interface{}, error) {
	var rsp, err  = x.RawRequest(fmt.Sprintf("%s/api/json", path));
	if err != nil { return nil, err; }
	var jobj map[string]interface{};
	err = json.Unmarshal([]byte(rsp), &jobj);
	return jobj, err;
}

