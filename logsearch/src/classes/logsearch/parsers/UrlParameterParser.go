package parsers
import (
    "fmt"
    "strings"
)

type UrlParameterParser struct {}
func (x UrlParameterParser) Parse(raw string)  (error, interface{}) {
    var rst = map[string]string{};
	var urlComps = strings.Split(raw, "?");
	if len(urlComps) < 2 {
		return nil, rst;
	}
	var parms = strings.Split(urlComps[1], "&");
	for _, pair := range parms {
		var pairComps = strings.Split(pair, "=");
		rst[pairComps[0]] = pairComps[1];
	}
    return nil, rst;
}

func (x UrlParameterParser) Unparse(obj interface{}) (error, string) {
	var rst = "?";
	var i = 0;
	var queryComps = obj.(map[string]string);
    for k, v := range queryComps {
    	if i > 0 { rst += "&"; }
        rst += fmt.Sprintf("%s=%s", k, v);
        i += 1;
	}
	return nil, rst;
}
