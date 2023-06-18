package esxi
import (
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
	"encoding/base64"
	"crypto/tls"
	"io/ioutil"
	"os"
	"math/rand"
	"strings"
	"time"
	"basgys/goxml2json"
	"esxi/actions"
)

type EsxiClient struct {
	sessionId             string
	baseEndpoint          string
	username              string
	pat                   string
}

func XmlToJson(raw string) (map[string]interface{}, string, error) {
    var rst = map[string]interface{}{};
    var err = json.Unmarshal([]byte(raw), &rst);
    if err != nil {
        var rawJson, err = goxml2json.Convert(strings.NewReader(raw));
        if err != nil {
            return nil, err.Error(), err;
        }
        err = json.Unmarshal([]byte(rawJson.String()), &rst);
        if err != nil {
            if raw != "" {
                rst = map[string]interface{}{
                    "data": raw,
                };
            } else {
                return nil, err.Error(), err;
            }
        }
        return rst, rawJson.String(), nil;
    }
    return rst, "", nil;
}

func NewEsxiClient(path string) (error, *EsxiClient) {
	var rst = &EsxiClient{};
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
	_, exists = jsonConfig["disableSsl"]; if exists {
		if jsonConfig["disableSsl"].(bool) {
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true};
		}
	}
	return nil, rst;
}

func (x EsxiClient) generateOperationId() string {
	var rst = "";
	var chars = strings.Split("a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z,0,1,2,3,4,5,6,7,8,9", ",");
	rand.Seed(time.Now().UnixNano());
    var maxLength = rand.Intn(4) + 4;
	for i := 0; i < maxLength; i++ {
		var randI = rand.Intn(len(chars) - 1);
		rst += chars[randI];
	}
	return rst;
}

func (x EsxiClient) buildEnvelope(request actions.BaseAction) string {
        var rst = `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/"`;
        rst += ` xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">`;
        rst += "<Header>";
        rst += fmt.Sprintf("<operationID>esxui-%s</operationID>", x.generateOperationId());
        rst += "</Header>";
        rst += "<Body>";
		rst += request.GenerateBody();
        rst += "</Body>";
        rst += "</Envelope>";
        return rst;
}

func (x *EsxiClient) authenticate() error {
	if x.sessionId != "" { return nil; }
	var client       = http.Client{};
	var body         = fmt.Sprintf(`<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<Header>
<operationID>esxui-%s</operationID>
</Header> 
<Body> 
<Login xmlns="urn:vim25"> 
<_this type="SessionManager">ha-sessionmgr</_this>
<userName>%s</userName> 
<password>%s</password>
<locale>en-US</locale>
</Login>
</Body>
</Envelope>`, x.generateOperationId(), x.username, x.pat);
	req, err         := http.NewRequest("POST", fmt.Sprintf("%s/sdk", x.baseEndpoint), bytes.NewBuffer([]byte(body)));
	if err != nil { return err; }
	req.Header.Add("Content-Type", "text/xml");
	rsp, err := client.Do(req);
	if err != nil { return err; }
	x.sessionId = strings.Replace(rsp.Cookies()[0].Value, `"`, "", -1);
	return nil;
}

func (x EsxiClient) RawRequest(headers map[string]string, request actions.BaseAction) (error, map[string]interface{}) {
	var err          = x.authenticate();
	if err != nil {
		return err, nil;
	}
	var client       = http.Client{};
	var endpoint     = fmt.Sprintf("%s/sdk", x.baseEndpoint); 
	req, err         := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(x.buildEnvelope(request))));
	if err != nil { return err, map[string]interface{}{}; }
	for h, v := range headers {
		req.Header.Add(h, v);
	}
	req.Header.Add("Content-Type", "text/xml");
	req.Header.Add("vmware-api-session-id", x.sessionId);
	rsp, err := client.Do(req);
	if err != nil { return nil, map[string]interface{}{}; }
	var rspData, err = ioutil.ReadAll(rsp.Body);
	if err != nil { return err, map[string]interface{}{}; }
	var jsonObj, err = XmlToJson(string(rspData));
	return err, jsonObj;
}

func (x EsxiClient) SessionId() string { return x.sessionId; }
func (x EsxiClient) BaseEndpoint() string { return x.baseEndpoint; }

func (x *EsxiClient) SetSessionId(a string) { x.sessionId = a; }
func (x *EsxiClient) SetBaseEndpoint(a string) { x.baseEndpoint = a; }
