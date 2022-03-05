package data
import (
	"testing"
	"os"
)

func initializeSQLite(config map[string]string) {
	config["connectionString"] = "test.db";
	connection = &DataConnectionSQLite{};
	connection.SetConnectionString(config["connectionString"]);
	connection.RunQuery("CREATE TABLE NAMEVALUE(NAME CHARACTER, VALUE CHARACTER)");
}

func cleanup() {
	os.Remove("test.db");
}

func TestSQLiteGetColumnNames(t *testing.T) {
	initializeSQLite(configuration);
	var var1 = connection.GetColumnNames("SELECT * FROM NAMEVALUE");
	if len(var1) == 0 {
		cleanup();
		t.Fail();
	}
	cleanup();
}

func TestSQLiteQuery(t *testing.T) {
	initializeSQLite(configuration);
	var var1 = connection.Query("SELECT * FROM NAMEVALUE");
	if var1 == nil {
		cleanup();
		t.Fail();
	}
	cleanup();
}

func TestSQLiteRunQuery(t *testing.T) {
	initializeSQLite(configuration);
	cleanup();
}

