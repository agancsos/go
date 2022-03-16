package api
import (
	"fmt"
	"net/http"
	"../helpers"
)

type restService struct {
	port                int
	serviceName         string
}

func NewRestService(port int, name string) *restService {
	var instance = &restService{};
	instance.port = port;
	instance.serviceName = name;
	return instance;
}


// API handlers
func (x restService) getVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf(`{"name":"%s", "version":"v. %s"}`, helpers.APPLICATION_NAME, helpers.APPLICATION_VERSION)));
}
/*****************************************************************************/


func (x *restService) StartServer() {
	http.HandleFunc("/api/version/", x.getVersion);

	http.ListenAndServe(fmt.Sprintf(":%d", x.port), nil);
}

