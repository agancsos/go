package data
import (
	"testing"
)

var configuration = map[string]string {
	"connectionString": "DRIVER=PSQL; DATABASE=postgres; HOST=localhost;",
	"username"        : "postgres",
	"password"        : "postgres",
};
var connection DataConnection;

func initializeOdbc(config map[string]string) {
	connection = &DataConnectionOdbc{};
	connection.SetUsername(config["username"]);
	connection.SetPassword(config["password"]);
	connection.SetConnectionString(config["connectionString"]);
}

func TestOdbcGetColumnNames(t *testing.T) {
	initializeOdbc(configuration);
	var var1 = connection.GetColumnNames("SELECT * FROM pg_tables");
	if len(var1) == 0 {
		t.Fail();
	}
}

func TestOdbcQuery(t *testing.T) {
	initializeOdbc(configuration);
	var var1 = connection.Query("SELECT * FROM pg_tables");
	if var1 == nil {
		t.Fail();
	}
}

func TestOdbcRunQuery(t *testing.T) {
	initializeOdbc(configuration);
}

