package common
import (
    "net/http"
	"bytes"
	"io/ioutil"
	"fmt"
	"errors"
	"strconv"
)

func InvokeGet(endpoint string, headers map[string]string) map[string]interface{} {
	var client = http.Client{};
	req, err := http.NewRequest("GET", endpoint, nil);
	for key, value := range headers {
		req.Header.Add(key, value);
	}
    rsp, err := client.Do(req);
    if err == nil {
		var temp, _ = strconv.Atoi(rsp.Status);
        if temp >= 300 {
            return map[string]interface{}{"error": errors.New(fmt.Sprintf("Received error: %d", temp)),};
        }
        rspData, _ := ioutil.ReadAll(rsp.Body);
		return StrToDictionary(rspData);
    }
    return nil;
}

func InvokePost(endpoint string, jsonBody map[string]string, headers map[string]string) map[string]interface{} {
    body := StrDictionaryToJsonString(jsonBody)
	var client = http.Client{};
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(fmt.Sprintf("%s", body))));
	for key, value := range headers {
        req.Header.Add(key, value);
    }
    rsp, err := client.Do(req);
    if err == nil {
		var temp, _ = strconv.Atoi(rsp.Status);
    	if temp >= 300 {
        	return map[string]interface{}{"error": errors.New(fmt.Sprintf("Received error: %d", temp)),};
    	}
        rspData, _ := ioutil.ReadAll(rsp.Body);
		return StrToDictionary(rspData);
    }
    return nil;
}

func EnsureRestMethod(a *http.Request, b string) (bool, string) {
	if a == nil || a.Method != b {
		return false, "";
	}
	var body, _ = ioutil.ReadAll(a.Body);
	if b == "POST" && string(body) == "" {
		return false, string(body);
	}
	return true, string(body);
}
