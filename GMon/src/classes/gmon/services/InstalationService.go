package services
import (
	"../../common"
	"../../sr"
	"../../data"
	"fmt"
	"io/ioutil"
)

type InstallationService struct {
	dbaUsername	   string
	dbaPassword	   string
	username	   string
	password	   string
	schema		   string
	server		   string
	driver		   string
	debug		   bool
	ss			   *SecurityService
	connection	   data.DataConnection;
}
var iis *InstallationService;

func GetInstallationServiceInstance() *InstallationService {
	if iis == nil {
		iis = &InstallationService{};
		iis.ss = GetSecurityServiceInstance();
	}
	return iis;
}

func (x *InstallationService) createSchema() {
	data, _ := ioutil.ReadFile(sr.GetSRInstance().GetDataConfigurationFilePath(""));
	var dbConfig = common.StrToDictionary(data);

	for _, p := range dbConfig["queries"].(map[string]interface{}) {
		if p.(map[string]string)["packageName"] == "create" {
			var sql = p.(map[string]string)["query"];
			var name = p.(map[string]string)["name"];
			println(sql);
			if sql != "" {
				if len(x.connection.Query(fmt.Sprintf("SELECT 1 FROM information_schema.TABLES WHERE TABLE_NAME ='%s' AND TABLE_SCHEMA = '%s'", name, x.schema)).Rows()) == 0 {
					if !x.debug {
						if !x.connection.RunQuery(sql) {
							panic(fmt.Sprintf("Failed to create table (%s)", name));
						}
					}
				}
			}
		}
	}
}

func (x *InstallationService) updateSchema() {
	data, _ := ioutil.ReadFile(sr.GetSRInstance().GetDataConfigurationFilePath(""));
	var dbConfig = common.StrToDictionary(data);

	for _, p := range dbConfig["queries"].(map[string]interface{}) {
		if p.(map[string]string)["packageName"] == "update" {
			var sql = p.(map[string]string)["query"];
			var name = p.(map[string]string)["name"];
			println(sql);
			if sql != "" {
				if len(x.connection.Query(fmt.Sprintf("SELECT 1 FROM information_schema.TABLES WHERE TABLE_NAME ='%s' AND TABLE_SCHEMA = '%s'", name, x.schema)).Rows()) > 0 {
					if !x.debug {
						if !x.connection.RunQuery(sql) {
							panic(fmt.Sprintf("Failed to update table (%s)", name));
						}
					}
				}
			}
		}
	}
}

func (x *InstallationService) Invoke() {
	x.connection = &data.DataConnectionOdbc{};
	x.connection.SetUsername(x.dbaUsername);
	x.connection.SetPassword(x.ss.GetDecoded(x.dbaPassword));
	x.connection.SetConnectionString(x.GetConnectionString());

	// Create user and container
	if !x.debug {
		if !x.connection.RunQuery(fmt.Sprintf("CREATE USER %s IDENTIFIED BY '%s'", x.username, x.ss.GetDecoded(x.password))) {
			println("User account already exists.  Skipping...");
		}
		if !x.connection.RunQuery(fmt.Sprintf("CREATE DATABASE %s", x.schema)) {
			panic("Failed to create database container...");
		}
		if !x.connection.RunQuery(fmt.Sprintf("GRANT ALL PRIVILEGES ON %s.* TO %s IDENTIFIED BY '%s'", x.schema, x.username, x.ss.GetDecoded(x.password))) {
			panic("Failed to grant permissions...");
		}
	}
	x.connection.SetUsername(x.username);
	x.connection.SetPassword(x.ss.GetDecoded(x.password));
	x.connection.SetConnectionString(x.connection.ConnectionString() + ";Database=" + x.schema);
	x.createSchema();
	x.updateSchema();

	// Save configuration file
	var configuration = "{";
	configuration += "\"version\":\"1.0.0\",";
	configuration += ("\"db\":{\"connectionString\":\"" + x.GetConnectionString() + "\",");
	configuration += ("\"schema\":\"" + x.schema + "\",\"username\":\"" + x.username + "\", \"password\":\"" + x.password + "\"}");
	configuration += "}";
	if !x.debug {
		(GetConfigurationServiceInstance()).SaveConfiguration(configuration);
	}
}

func (x *InstallationService) SetDbaUsername(a string) {
	x.dbaUsername = a;
	x.connection.SetUsername(a);
}

func (x *InstallationService) SetDbaPassword(a string) {
	x.dbaPassword = a;
	x.connection.SetPassword(x.ss.GetDecoded(a));
}

func (x *InstallationService) GetConnectionString() string {
	var result = "DRIVER=";
	if x.driver != "" {
		result += x.driver;
	} else {
		result += "MySQL UNICODE DRIVER";
	}
	result += ";";
	result += "SERVER=";
	result += x.server;
	return result;
}

func (x *InstallationService) SetUsername(a string) { x.username = a; }
func (x *InstallationService) SetPassword(a string) { x.password = a; }
func (x *InstallationService) SetSchema(a string) { x.schema = a; }
func (x *InstallationService) SetServer(a string) { x.server = a; }
func (x *InstallationService) SetDriver(a string) { x.driver = a; }
func (x *InstallationService) SetDebug(a bool) { x.debug = a; }
func (x *InstallationService) GetDbaUsername() string { return x.dbaUsername; }
func (x *InstallationService) GetDbaPassword() string { return x.dbaPassword; }
func (x *InstallationService) GetUsername() string { return x.username; }
func (x *InstallationService) GetPassword() string { return x.password; }
func (x *InstallationService) GetSchema() string { return x.schema; }
func (x *InstallationService) GetServer() string { return x.server; }
func (x *InstallationService) GetDriver() string { return x.driver; }
func (x *InstallationService) GetDebug() bool { return x.debug;}

