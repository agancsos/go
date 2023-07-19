package hercules
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"net/http"
	"regexp"
)

type HerculesClient struct {
	BaseEndpoint                 string   `json:baseEndpoint`
}

func DictionaryToJsonString (a map[string]interface{}) string {
    var result = "";
    var i = 0;
    for key, value := range a {
        if i > 0 { result += "&"; }
        result += fmt.Sprintf("%s=%s", key, url.QueryEscape(fmt.Sprintf("%v", value)));
        i += 1;
    }
    result += "";
    return result;
}

func NewClient(baseEndpoint string) (*HerculesClient, error) {
	var client =  &HerculesClient{BaseEndpoint: baseEndpoint};
	return client, nil;
}

func parseResponse(raw string) map[string]interface{} {
	if !strings.Contains(raw, "<table>") {
		rex, err := regexp.Compile("(?s)<PRE>.*</PRE>");
		if err != nil {
			return map[string]interface{} {
				"raw": strings.Replace(raw, "\n", "\\n", -1),
				"error": err.Error(),
			};
		}
		var match = rex.FindAllString(raw, -1)[0];
		return map[string]interface{} {
			"raw": strings.Replace(raw, "\n", "\\n", -1),
			"data": strings.Replace(strings.Replace(strings.Replace(match, "\n", "\\n", -1), "<PRE>", "", -1), "</PRE>", "", -1),
		}
	} else {
		var tables = []interface{}{};
		rex1, err := regexp.Compile("(?s)<table>.*?</table>");
		if err != nil {
			return map[string]interface{} {
				"raw": strings.Replace(raw, "\n", "\\n", -1),
				"error": err.Error(),
			}
		}
		var tables2 = rex1.FindAllString(raw, -1);
		for _, table := range tables2 {
			var tableData = []interface{}{};
			rex2, err := regexp.Compile("(?s)<th>.*?</th>");
			if err != nil {
				return map[string]interface{} {
					"raw": strings.Replace(raw, "\n", "\\n", -1),
					"error": err.Error(),
				}
			}
			var fields = []string{};
			var rawData = rex2.FindAllString(table, -1);
			for _, field := range rawData {
				fields = append(fields, strings.Replace(strings.Replace(field, "<th>", "", -1), "</th>", "", -1));
			}
			rex2, err = regexp.Compile("(?s)<tr>.*?</tr>");
			if err != nil {
				return map[string]interface{} {
					"raw": strings.Replace(raw, "\n", "\\n", -1),
					"error": err.Error(),
				}
			}
			var lines = rex2.FindAllString(table, -1);
			for _, line := range lines {
				rex3, err := regexp.Compile("(?s)<td>.*?</td>");
				if err != nil {
					return map[string]interface{} {
						"raw": strings.Replace(raw, "\n", "\\n", -1),
						"error": err.Error(),
					}
				}
				var rawItem = map[string]interface{}{};
				var values  = rex3.FindAllString(line, -1);
				for i := range values {
					rawItem[fields[i]] = strings.Replace(strings.Replace(values[i], "<td>", "", -1), "</td>", "", -1);
				}
				tableData = append(tableData, rawItem);
			}
			tables = append(tables, tableData);
		}
		return map[string]interface{} {
			"raw": strings.Replace(raw, "\n", "\\n", -1),
			"data": tables,
		}				
	}
}

func (x HerculesClient) RawRequest(path string, headers map[string]string) (map[string]interface{}, error) {
	var client       = http.Client{};
	if !strings.Contains(path, "http") { path = fmt.Sprintf("%s/%s", x.BaseEndpoint, path); }
	req, err         := http.NewRequest("GET", path, nil);
	if err != nil { return nil, err; }
	for h, v := range headers {
		req.Header.Add(h, v);
	}
	rsp, err := client.Do(req);
	if err != nil { return nil, err; }
	rspData, err := ioutil.ReadAll(rsp.Body);
	if err != nil { return nil, err; }
	return parseResponse(string(rspData)), nil;
}

func (x HerculesClient) RawPostRequest(path string, headers map[string]string, body map[string]interface{}) (map[string]interface{}, error) {
	var client	    = http.Client{};
	if !strings.Contains(path, "http") { path = fmt.Sprintf("%s/%s", x.BaseEndpoint, path); }
	req, err		:= http.NewRequest("POST", path, bytes.NewBuffer([]byte(DictionaryToJsonString(body))));
	for h, v := range headers {
		req.Header.Add(h, v);
	}
	rsp, err		:= client.Do(req);
	if err != nil { return nil, err; }
	rspData, err	:= ioutil.ReadAll(rsp.Body);
	if err != nil { return nil, err; }
	return parseResponse(string(rspData)), nil;
}

