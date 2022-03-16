package data
import (
	"strings"
)
var SupportedProviders = map[string]string {
	"sqlite:" : "SQLite",
	"odbc:"   : "ODBC",
}


func CreateConnection(x string, a string, b string) DataConnection {
	if strings.Replace(x, "sqlite:", "", -1) != x {
		return &DataConnectionSQLite{path: strings.Replace(x, "sqlite:", "", -1)};
	} else if strings.Replace(x, "odbc:", "", -1) != x {
		return &DataConnectionOdbc{DatabaseConnectionString: strings.Replace(x, "odbc:", "", -1), DatabaseUsername: a, DatabasePassword: b};
	}
	return &EmptyConnection{};
}

// DataColumn
type DataColumn struct {
	columnName  string
	columnValue string
	columnType  string
}
func (x DataColumn) Name() string { return x.columnName; }
func (x DataColumn) Value() string { return x.columnValue; }
func (x DataColumn) Type() string { return x.columnType; }
func (x *DataColumn) SetName(y string) { x.columnName = y; }
func (x *DataColumn) SetValue(y string) { x.columnValue = y; }
func (x *DataColumn) SetType(y string) { x.columnType = y; }
/*********************************/

// DataRow
type DataRow struct {
	columns []*DataColumn
}
func (x DataRow) Columns() []*DataColumn { return x.columns; }
func (x DataRow) Column(name string) *DataColumn {
	for _, column := range x.columns {
		if strings.EqualFold(column.Name(), name) {
			return column;
		}
	}
	return &DataColumn{};
}
func (x DataRow) Contains(name string) bool {
	for i := range x.Columns() {
        if x.Columns()[i].Name() == name {
            return true;
		}
	}
	return false;
}
func (x *DataRow) AddColumn(y *DataColumn) {
	if !x.Contains(y.Name()) {
		x.columns = append(x.columns, y);
	}
}
/*********************************/

// DataTable
type DataTable struct {
    rows []*DataRow
}
func (x DataTable) Rows() []*DataRow { return x.rows; }
func (x *DataTable) AddRow(y *DataRow) {
	x.rows = append(x.rows, y);
}
/*********************************/

// IDataConnection
type DataConnection interface {
	ConnectionString()    string
	Username()            string
	Password()            string
	Query(a string)          *DataTable
	GetTableNames()          []string
	GetColumnNames(a string) []string
    RunQuery(a string)        bool
	SetUsername(a string)
	SetPassword(a string)
	SetConnectionString(a string)
	connect()                bool
	disconnect()
}
/*********************************/


// EmptyConnection
type EmptyConnection struct { }
func (x *EmptyConnection) Query(a string) *DataTable { return &DataTable{}; }
func (x *EmptyConnection) RunQuery(a string) bool { return true; }
func (x EmptyConnection) GetColumnNames(a string) []string { var result []string; return result;  }
func (x *EmptyConnection) ConnectionString() string { return "" }
func (x EmptyConnection) Username() string { return ""; }
func (x EmptyConnection) Password() string { return ""; }
func (x EmptyConnection) GetTableNames() []string { var result []string; return result; }
func (x *EmptyConnection) SetConnectionString(a string) {}
func (x *EmptyConnection) SetUsername(a string) {}
func (x *EmptyConnection) SetPassword(a string) {}
func (x *EmptyConnection) connect() bool { return true; }
func (x *EmptyConnection) disconnect() {}
/*********************************/

