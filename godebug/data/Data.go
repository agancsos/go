package data
import (
	"strings"
)
var SupportedProviders = map[string]string {
	"sqlite:" : "SQLite",
}


func CreateConnection(x string, a string, b string) IDataConnection {
	if strings.Replace(x, "sqlite:", "", -1) != x {
		return &SQLiteConnection{path: strings.Replace(x, "sqlite:", "", -1)};
	}
	return &EmptyConnection{};
}

// DataColumn
type DataColumn struct {
	columnName  string
	columnValue string
	columnType  string
}
func (x DataColumn) GetName() string { return x.columnName; }
func (x DataColumn) GetValue() string { return x.columnValue; }
func (x DataColumn) GetType() string { return x.columnType; }
func (x DataColumn) SetName(y string) { x.columnName = y; }
func (x DataColumn) SetValue(y string) { x.columnValue = y; }
func (x DataColumn) SetType(y string) { x.columnType = y; }
/******************************************************************/

// DataRow
type DataRow struct {
	columns []DataColumn
}
func (x DataRow) GetColumns() []DataColumn { return x.columns; }
func (x DataRow) GetColumn(name string) DataColumn {
	for i := 0; i < len(x.columns); i++ {
		if x.columns[i].GetName() == name {
			return x.columns[i];
		}
	}
	return DataColumn{};
}
func (x DataRow) Contains(name string) bool {
	for i := 0; i < len(x.columns); i++ {
        if x.columns[i].GetName() == name {
            return true;
		}
	}
	return false;
}
func (x DataRow) AddColumn(y DataColumn) {
	if !x.Contains(y.GetName()) {
		x.columns = append(x.columns, y);
	}
}
/******************************************************************/

// DataTable
type DataTable struct {
    rows []DataRow
}
func (x DataTable) GetRows() []DataRow { return x.rows; }
func (x DataTable) AddRow(y DataRow) { x.rows = append(x.rows, y); }
/******************************************************************/

// IDataConnection
type IDataConnection interface {
	GetConnectionString()    string
	GetUsername()            string
	GetPassword()            string
	Query(a string)          DataTable
	GetTableNames()          []string
	GetColumnNames(a string) []string
    RunQuery(a string)        bool
	SetUsername(a string)
	SetPassword(a string)
	connect()                bool
	disconnect()
}
/******************************************************************/


// EmptyConnection
type EmptyConnection struct {
}
func (x *EmptyConnection) Query(a string) DataTable { return DataTable{}; }
func (x *EmptyConnection) RunQuery(a string) bool { return true; }
func (x *EmptyConnection) GetColumnNames(a string) []string { var result []string; return result;  }
func (x *EmptyConnection) GetConnectionString() string { return "" }
func (x *EmptyConnection) GetUsername() string { return ""; }
func (x *EmptyConnection) GetPassword() string { return ""; }
func (x *EmptyConnection) GetTableNames() []string { var result []string; return result; }
func (x *EmptyConnection) SetUsername(a string) {}
func (x *EmptyConnection) SetPassword(a string) {}
func (x *EmptyConnection) connect() bool { return true; }
func (x *EmptyConnection) disconnect() {}
/******************************************************************/

