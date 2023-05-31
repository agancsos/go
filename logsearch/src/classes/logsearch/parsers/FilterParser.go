package parsers
import (
	"fmt"
	"strings"
)

type FilterParser struct {}
func (x FilterParser) Parse(raw string)  (error, interface{}) {
	var rst = map[string]string{};
	if strings.Contains(strings.ToUpper(raw), " WHERE ") {
        var filterString = strings.Split(raw, " WHERE ")[1];
        var filterPairs  = strings.Split(filterString, " AND ");
		if len(filterPairs) < 1 {
			filterPairs  = strings.Split(filterString, " OR ");
		}
        for _, pairs := range filterPairs {
			for _, char := range []string{">", ">=", "<", "<=", "LIKE",} {
				if strings.Contains(strings.ToUpper(pairs), char) {
					var pair = strings.Split(pairs, char);
					rst[pair[0]] = pair[1];
				}
			}
        }
    } else {
		var filterPairs = strings.Split(raw, "&");
		for _, pairs := range filterPairs {
			var pair = strings.Split(pairs, "=");
			if len(pair) > 1 {
				rst[pair[0]] = pair[1];
			}
		}
    }	
	return nil, rst;
}

func (x FilterParser) Unparse(obj interface{}) (error, string) {
	var filter = obj.(map[string]string);
	var rst = "SELECT * FROM LOGS";
	if len(filter) > 0 {
		rst += " WHERE ";
		var i = 0;
		for k, v := range filter {
			if i > 0 { rst += " AND "; }
			rst += fmt.Sprintf(`%s LIKE "%\%s%"`, k, v);
		}
	}
	return nil, rst;
}

