package main
import (
    "os"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "bytes"
)

type RestHelper struct {
    BasePath string
}

func (a *RestHelper) InvokeGet(operation string) map[string]string {
    rsp, err := http.Get(a.BasePath + "/" + operation);
    if err == nil {
        rspData, _ := ioutil.ReadAll(rsp.Body);
        var obj map[string]string;
        json.Unmarshal([]byte(rspData), &obj);
        return obj;
    }
    return nil;
}

func (a *RestHelper) InvokePost(operation string, jsonBody map[string]string) map[string]string {
    body, _ := json.Marshal(jsonBody)
    rsp, err := http.Post(a.BasePath + "/" + operation, "application/json", bytes.NewBuffer(body));
    if err == nil {
        rspData, _ := ioutil.ReadAll(rsp.Body);
        var obj map[string]string;
        json.Unmarshal([]byte(rspData), &obj);
        return obj;
    }
    return nil;
}

func main() {
    var helper = &RestHelper{BasePath: ""};
    var rsp = helper.InvokeGet("");
    println(rsp[""]);
    os.Exit(0);
}
