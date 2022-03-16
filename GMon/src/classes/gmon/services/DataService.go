package services
import (
	"../../data"
	"fmt"
	"../models"
)

type DataService struct {
	connection  data.DataConnection
	ss          *SecurityService
}
var __data_service__ *DataService;

func (x *DataService) Initialize() {
	x.ss = GetSecurityServiceInstance();
	x.connection = &data.DataConnectionOdbc{};
	var dbNode = (GetConfigurationServiceInstance()).GetKey("database", "").(map[string]interface{});
	x.connection.SetUsername(dbNode["username"].(string));
	x.connection.SetPassword(x.ss.GetDecoded(dbNode["password"].(string)));
	x.connection.SetConnectionString(x.buildConnectionString());
}

func GetDataServiceInstance() IService {
    if __data_service__ == nil {
        __data_service__ = &DataService{};
		__data_service__.Initialize();
    }
    return __data_service__;
}

func (x *DataService) buildConnectionString() string {
	var result = "DRIVER=";
    if (GetConfigurationServiceInstance()).GetKey("database", "") != "" {
        var dbNode = (GetConfigurationServiceInstance()).GetKey("database", "").(map[string]interface{});
        if dbNode["driver"] != nil && dbNode["driver"].(string) != "" {
            result += dbNode["driver"].(string);
        } else {
            result += "MySQL UNICODE ODBC";
		}
        result += (";Database=" + dbNode["database"].(string) + "");
        result += ";";
        result += "SERVER=";
        result += dbNode["serverName"].(string);
        result += ";";
	}
    return result;
}

func (x *DataService) RunServiceQuery(query string) bool {
	return x.connection.RunQuery(query);
}

func (x *DataService) ServiceQuery(query string) *data.DataTable {
	return x.connection.Query(query);
}

func (x *DataService) GetSingletVerification(hostname string) bool {
	var result = x.ServiceQuery("SELECT 1 FROM AGENTS WHERE AGENTSTATE IN (" +
               fmt.Sprintf("'%v',", models.NMS_INITIALIZED) +
               fmt.Sprintf("'%v',", models.NMS_INITIALIZING) +
               fmt.Sprintf("'%v')", models.NMS_STOPPING));
    if len(result.Rows()) < 1 {
		return true;
	}
	return false;
}

