package common
import (
    "io/ioutil"
    "net/http"
    "encoding/json"
	"bytes"
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

func (a *RestHelper) InvokePost(endpoint string, jsonBody map[string]string,  headers map[string]string) map[string]interface{} {
    body := StrDictionaryToJsonString(jsonBody)
	var client = http.Client{};
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader([]byte(body)));
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

func StrToDictionary(s []byte) map[string]interface{} {
	var obj map[string]interface{};
	json.Unmarshal(s, &obj);
	return obj;
}

func DictionaryToJsonString (a map[string]interface{}) string {
	var result = "{";
	for key, value := range a {
		result += fmt.Sprintf("\"%s\":\"%v\"", key, value);
	}
	result += "}";
	return result;
}

func StrDictionaryToJsonString (a map[string]string) string {
    var result = "{";
    for key, value := range a {
        result += fmt.Sprintf("\"%s\":\"%s\"", key, value);
    }
    result += "}";
    return result;
}
