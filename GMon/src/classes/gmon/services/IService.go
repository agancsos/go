package services
import (
	"../../common"
	"../../sr"
	"io/ioutil"
	"net/http"
	"fmt"
)

// Service interface
type IService interface {
	Initialize()
}
/*****************************************************************************/

// Default service
type EmptyService struct{
	instance *EmptyService
}
func (x *EmptyService) Initialize() { }
/*****************************************************************************/

// ServiceProvider interface
type IServiceProvider interface {
	RegisterService (a string, b IService);
	GetService (a string) IService;
	ContainsService (a string) bool
	GetServices() map[string]IService;
}
/*****************************************************************************/

// Default service provider
type EmptyServiceProvider struct {
	instance *EmptyServiceProvider;
	services map[string]IService;
}
func (x *EmptyServiceProvider) RegisterService (a string, b IService) {}
func (x *EmptyServiceProvider) GetService (a string) IService { return nil; }
func (x *EmptyServiceProvider) ContainsService(a string) bool { return true; }
func (x *EmptyServiceProvider) GetServices() map[string]IService { return x.services; }
/*****************************************************************************/

// ConfigurationService
type ConfigurationService struct {dictionary    map[string]interface{}}
var __configuration_service__ *ConfigurationService;
func GetConfigurationServiceInstance() *ConfigurationService {
	if __configuration_service__ == nil {
		__configuration_service__ = &ConfigurationService{};
	}
	return __configuration_service__;
}
func (x *ConfigurationService) GetProperty(name string, defaultValue string) interface{} {
    rawXml, _ := ioutil.ReadFile((sr.GetSRInstance()).GetConfigurationFile());
    x.dictionary = common.StrToDictionary(rawXml);
    if x.dictionary[name] == nil { return defaultValue; }
    return x.dictionary[name];
}
func (x *ConfigurationService) SaveConfiguration(raw string) {
    ioutil.WriteFile((sr.GetSRInstance()).GetConfigurationFile(), []byte(raw), 'w');
}
func (x *ConfigurationService) GetKey(name string, defaultValue string) interface{} {
	return x.GetProperty(name, defaultValue);
}
/*****************************************************************************/


// Service helpers
func EnsureAuthenticated(w http.ResponseWriter, r *http.Request, m string) (bool, string) {
	okay, data := common.EnsureRestMethod(r, m);
    if !okay {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"Invalid method\"}")));
        return false, data;
    }
    if !(GetLocalAuthService()).IsValid(sr.ExtractApiToken(r), "") {
        w.Write([]byte(fmt.Sprintf("{\"message\":\"User is invalid\"}")));
        return false, data;
    }
    return true, data;
}
/*****************************************************************************/

// Template
// Interface
/*****************************************************************************/

// Local service
/*****************************************************************************/

// Remote service
/*****************************************************************************/
