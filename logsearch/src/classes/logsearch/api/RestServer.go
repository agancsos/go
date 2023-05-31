package api
import (
    "net/http"
	"../../common"
	"../services"
	"../parsers"
    "fmt"
)

type restServer struct {
    port             int
	lls              *services.LogSearchService
}

func NewRestServer(port int) *restServer {
    var rst = &restServer{port: port};
	rst.lls = services.GetLogSearchServiceInstance();
    return rst;
}

func (x restServer) test(w http.ResponseWriter, r *http.Request) {
	okay, _ := common.EnsureRestMethod(r, "GET");
	if !okay{ return; }
	w.Write([]byte(`{"result": "HELLO MOTTO"}`));
}	

func (x restServer) purge(w http.ResponseWriter, r *http.Request) {
    okay, _ := common.EnsureRestMethod(r, "POST");
    if !okay{ return; }
	var err = x.lls.PurgeLogs();
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"result": "%s"}`, err.Error())));
		return;
	} else {
    	w.Write([]byte(`{"result": "SUCCESS"}`));
	}
}  

func (x restServer) query(w http.ResponseWriter, r *http.Request) {
    okay, _ := common.EnsureRestMethod(r, "GET");
    if !okay{ return; }
	var _, filters = (&parsers.UrlParameterParser{}).Parse(r.URL.String());
	var rst = "{";
	rst += `"result": [`;
	var err, events = x.lls.QueryLogs(filters.(map[string]string));
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"result": "%s"}`, err.Error())));
		return;
	}
	for i, event := range events {
		if i > 0 { rst += ","; }
		rst += event.ToJsonString();
	}
	rst += "]";
	rst += "}";
    w.Write([]byte(rst));
}  

func (x restServer) addEvent(w http.ResponseWriter, r *http.Request) {
    okay, data := common.EnsureRestMethod(r, "POST");
    if !okay{ return; }
	var dataJson = common.StrToStrDictionary(data);
	var err = x.lls.AddEvent(dataJson);
    if err != nil {
        w.Write([]byte(fmt.Sprintf(`{"result": "%s"}`, err.Error())));
		return;
    } else {
        w.Write([]byte(`{"result": "SUCCESS"}`));
    }
}

func (x *restServer) StartServer() {
    http.HandleFunc("/test/", x.test);
    http.HandleFunc("/purge/", x.purge);
    http.HandleFunc("/query/", x.query);
	http.HandleFunc("/add/", x.addEvent);
    http.ListenAndServe(fmt.Sprintf(":%d", x.port), nil);
}
