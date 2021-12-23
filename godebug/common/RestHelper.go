package common
import (
    "net/http"
	"bytes"
	"io/ioutil"
	"fmt"
)

type RestHelper struct {
	BasePath   string
}

func (a *RestHelper) InvokeGet(endpoint string, headers map[string]string) map[string]interface{} {
	var client = http.Client{};
	req, err := http.NewRequest("GET", endpoint, nil);
	for key, value := range headers {
		req.Header.Add(key, value);
	}
    rsp, err := client.Do(req);
    if err == nil {
        rspData, _ := ioutil.ReadAll(rsp.Body);
		return StrToDictionary(rspData);
    }
    return nil;
}

func (a *RestHelper) InvokePost(endpoint string, jsonBody map[string]string, headers map[string]string) map[string]interface{} {
    body := StrDictionaryToJsonString(jsonBody)
	var client = http.Client{};
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(fmt.Sprintf("%s", body))));
	for key, value := range headers {
        req.Header.Add(key, value);
    }
    rsp, err := client.Do(req);
    if err == nil {
        rspData, _ := ioutil.ReadAll(rsp.Body);
		return StrToDictionary(rspData);
    }
    return nil;
}

<<<<<<< HEAD
func EnsureRestMethod(a *http.Request, b string) (bool, string) {
	if a == nil || a.Method != b {
		return false, "";
	}
	var body, _ = ioutil.ReadAll(a.Body);
	if b == "POST" && string(body) == "" {
		return false, string(body);
	}
	return true, string(body);
=======
func EnsureRestMethod(a *http.Request, b string) bool {
	if a == nil || a.Method != b {
		return false;
	}
	var body, _ = ioutil.ReadAll(a.Body);
	if b == "POST" && string(body) == "" {
		return false;
	}
	return true;
>>>>>>> 0e1c84c30b02a7c14fc2c37c72407b4c56c57873
}
