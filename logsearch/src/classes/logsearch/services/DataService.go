package services
import (
	"../../data"
	"fmt"
	"os"
	"path/filepath"
)

type DataService struct {
	basePath            string
}
var __data_service__ *DataService;

func (x *DataService) Initialize() {
	var queries = []string {
		"CREATE TABLE IF NOT EXISTS FLAGS (FLAG_NAME CHARACTER NOT NULL, FLAG_VALUE CHARACTER DEFAULT '', LAST_UPDATED_DATE TIMESTAMP DEFAULT CURRENT_TIMESTAMP)",
		"CREATE TABLE IF NOT EXISTS MESSAGES (MESSAGE_ID INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, MESSAGE_LEVEL INTEGER DEFAULT '9', MESSAGE_CATEGORY INTEGER DEFAULT '1', MESSAGE_TEXT CHARACTER DEFAULT '', LAST_UPDATED_DATE TIMESTAMP DEFAULT CURRENT_TIMESTAMP)",
		"CREATE TABLE IF NOT EXISTS LOGS (LOG_ID INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, LOG_SOURCE CHARACTER, LOG_TEXT CHARACTER , LAST_UPDATED_DATE TIMESTAMP DEFAULT CURRENT_TIMESTAMP)",
	};
	for _, query := range queries {
		x.RunServiceQuery(query);
	}
}

func (x DataService) newConnection() data.DataConnection {
	var result data.DataConnection;
	result = &data.DataConnectionSQLite{};
	result.SetConnectionString(fmt.Sprintf("%s/logsearch.sqlite", x.basePath));
	return result;
}

func GetDataServiceInstance() *DataService {
	if __data_service__ == nil {
		__data_service__ = &DataService{};
		var binaryPath, _ = filepath.Abs(filepath.Dir(os.Args[0]));
		__data_service__.basePath = binaryPath;
		__data_service__.Initialize();
    }
    return __data_service__;
}

func (x DataService) RunServiceQuery(query string) bool {
	return x.newConnection().RunQuery(query);
}

func (x DataService) ServiceQuery(query string) *data.DataTable {
	conn     := x.newConnection();
	rst, err := conn.Query(query);
	if err != nil {
		return &data.DataTable{};
	}
	return rst;
}

