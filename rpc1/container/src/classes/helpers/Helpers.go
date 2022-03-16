package helpers
import (
	"../common"
	"io/ioutil"
	"../sr"
)

// ConfigurationService
type ConfigurationService struct {
    dictionary       map[string]interface{}
}
var csi *ConfigurationService;
func GetConfigurationServiceInstance() *ConfigurationService {
    if csi == nil {
        csi = &ConfigurationService{};
    }
    return csi;
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
