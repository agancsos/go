package services
import (
	"fmt"
	"strconv"
	"net/http"
	"../../sr"
)

type RestService struct {
	servicePort    int
	serviceName    string
	dbts           *DbTraceService
}
var rs *RestService;
func GetRestServiceInstance() *RestService {
	if rs == nil {
		rs = &RestService{}
		var cs = GetConfigurationServiceInstance();
		rs.servicePort = 4445;
		rs.servicePort, _ = strconv.Atoi(cs.GetProperty("apiPort", "4445").(string));
		rs.dbts = GetDbTraceServiceInstance().(*DbTraceService);
		rs.serviceName = "Gancsos Monitor REST Service";
	}
	return rs;
}

func (x *RestService) GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("{\"version\":\"%s\"}", sr.APPLICATION_VERSION)));
}

func (x *RestService) Initialize() {
	// Call Initialize on all remote services to start the endpoints
	GetRestAuthService().Initialize();
	GetRestRoleServiceInstance().Initialize();

	//Configure common endpoints
	http.HandleFunc("/version/", x.GetVersion);

	// Start listener
	http.ListenAndServe(fmt.Sprintf(":%d", x.servicePort), nil);
}

