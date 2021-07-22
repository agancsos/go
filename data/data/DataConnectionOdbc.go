package data

/*
#include <sql.h>
#include <sqlext.h>
#cgo linux LDFLAGS: -lodbc
#cgo darwin LDFLAGS: -liodbc
#cgo linux CFLAGS: -I"/usr/local/include/" -std=c11
#cgo darwin CFLAGS: -I"/usr/local/include/" -std=c11
*/
import "C";
import (
	"../common"
	"unsafe"
	"fmt"
)
var MAX_COLUMN_LENGTH = 1024;

type DataConnectionOdbc struct {
	hEnv             C.SQLHENV
	hHandler         C.SQLHDBC
	ConnectionString string
	Username         string
	Password         string
	IsConnected      bool
}
func (x *DataConnectionOdbc) check(a C.SQLRETURN, b string) {
	if a != C.SQL_SUCCESS && a != C.SQL_SUCCESS_WITH_INFO && a != 100  && b != "Get data" {
		panic(fmt.Sprintf("%s (%v)", b, a));
	}
}
func (x *DataConnectionOdbc) Query(a string) *DataTable {
	var result = &DataTable{};
	var statement C.SQLHSTMT;
    var columns = x.GetColumnNames(a);
    if x.connect() {
        x.check(C.SQLAllocHandle(C.SQL_HANDLE_STMT, x.hHandler, &statement), "Allocate statement (Query)");
        x.check(C.SQLExecDirect(statement, (*C.uchar)(common.ToConstStr(a)), C.SQL_NTS), "Execute");
        var row = -1;
        var status C.SQLRETURN;
        for ;; {
            row++;
            status = C.SQLFetch(statement);
            if status != C.SQL_SUCCESS && status != C.SQL_SUCCESS_WITH_INFO {
                if status != 100 {
					println(fmt.Sprintf("Error occurred on fetch... %v", status));
				}
                break;
            }
            var tempRow = &DataRow{};
            for col := 0; col < len(columns); col++ {
                var data  *C.SQLCHAR;
                var dataLen C.SQLLEN;
                var tempColumn = &DataColumn{};
                tempColumn.SetName(columns[col]);
                x.check(C.SQLGetData(statement, (C.ushort)(col + 1), C.SQL_C_CHAR, (C.SQLPOINTER)(data), (C.long)(MAX_COLUMN_LENGTH), &dataLen), "Get data");
                var dataValue = C.GoString((*C.char)(unsafe.Pointer(data)))
                tempColumn.SetValue(dataValue);
                tempRow.AddColumn(tempColumn);
            }
            result.AddRow(tempRow);
        }
        if statement != nil {
            x.check(C.SQLFreeHandle(C.SQL_HANDLE_STMT, statement), "Free statement");
        }
        x.disconnect();
    }
	return result;
}
func (x *DataConnectionOdbc) RunQuery(a string) bool {
	var result = false;
    var statement C.SQLHSTMT;
    if x.connect() {
        x.check(C.SQLAllocHandle(C.SQL_HANDLE_STMT, x.hHandler, &statement), "Allocate statement (RunQuery)");
        _ = C.SQLExecDirect(statement, (*C.SQLCHAR)(common.ToConstStr(a)), C.SQL_NTS);
		_ = C.SQLExecDirect(statement, (*C.SQLCHAR)(common.ToConstStr("COMMIT")), C.SQL_NTS);
        result = true;
        if statement != nil {
            x.check(C.SQLFreeHandle(C.SQL_HANDLE_STMT, statement), "Free statement");
        }
		x.check(C.SQLAllocHandle(C.SQL_HANDLE_STMT, x.hHandler, &statement), "Allocate statement (RunQuery)");
		if statement != nil {
            x.check(C.SQLFreeHandle(C.SQL_HANDLE_STMT, statement), "Free statement");
        }
        x.disconnect();
	}
    return result;
}
func (x *DataConnectionOdbc) GetColumnNames(a string) []string {
	var result []string;
	var statement C.SQLHSTMT;
    var cols C.SQLSMALLINT;
    if x.connect() {
        x.check(C.SQLAllocHandle(C.SQL_HANDLE_STMT, x.hHandler, &statement), "Allocate statement (GetColumns)");
        x.check(C.SQLExecDirect(statement, (*C.uchar)(common.ToConstStr(a)), C.SQL_NTS), "Execute");
        x.check(C.SQLNumResultCols(statement, &cols), "Column count");
        for i := 0; i < int(cols); i++ {
			var columnName *C.SQLCHAR;
            var columnNameLen C.SQLSMALLINT;
            x.check(C.SQLColAttribute(statement, (C.ushort)(i + 1), C.SQL_DESC_LABEL, (C.SQLPOINTER)(columnName),
				(C.short)((len(C.GoString((*C.char)(unsafe.Pointer(columnName)))) * 5)), &columnNameLen, nil), "Column attribute");
            var value = C.GoString((*C.char)(unsafe.Pointer(columnName)));
            result = append(result, value);
        }
        if statement != nil {
            x.check(C.SQLFreeHandle(C.SQL_HANDLE_STMT, statement), "Release statement");
        }
        x.disconnect();
    }
	return result;
}
func (x *DataConnectionOdbc) GetTableNames() []string {
	var result []string;
	return result;
}
func (x *DataConnectionOdbc) connect() bool {
	if x.IsConnected { return true; }
	x.check(C.SQLAllocHandle(C.SQL_HANDLE_ENV, nil, &x.hEnv), "Allocation of environment");
    x.check(C.SQLSetEnvAttr(x.hEnv, C.SQL_ATTR_ODBC_VERSION, C.SQLPOINTER(uintptr(C.SQL_OV_ODBC2)), 0), "Set version");
    x.check(C.SQLAllocHandle(C.SQL_HANDLE_DBC, x.hEnv, &x.hHandler), "Allocation of connection");
	x.check(C.SQLDriverConnect(x.hHandler, nil, (*C.SQLCHAR)((*C.uchar)(common.ToConstStr(x.ConnectionString))), C.SQL_NTS, nil, 0, nil, C.SQL_DRIVER_COMPLETE), "Connect");
    C.SQLSetConnectAttr(x.hHandler, C.SQL_ATTR_AUTOCOMMIT, C.SQLPOINTER(uintptr(1)), 0);
	x.IsConnected = true;
	return true;
}
func (x *DataConnectionOdbc) disconnect() {
	x.check(C.SQLDisconnect(x.hHandler), "Disconnect");
	x.check(C.SQLFreeHandle(C.SQL_HANDLE_DBC, x.hHandler), "Release connection");
    x.check(C.SQLFreeHandle(C.SQL_HANDLE_ENV, x.hEnv), "Release environment");
    x.IsConnected = false;
}
func (x *DataConnectionOdbc) GetConnectionString() string { return x.ConnectionString; }
func (x *DataConnectionOdbc) GetUsername() string { return x.Username; }
func (x *DataConnectionOdbc) GetPassword() string { return x.Password; }
func (x *DataConnectionOdbc) SetUsername(a string) { x.Username = a; }
func (x *DataConnectionOdbc) SetPassword(a string) { x.Password = a; }
func (x *DataConnectionOdbc) SetConnectionString(a string) {
	x.ConnectionString = a;
	if string(x.ConnectionString[len(x.ConnectionString) - 1]) != ";" { x.ConnectionString += ";"; }
    x.ConnectionString += fmt.Sprintf("UID=%s;PWD=%s;", x.Username, x.Password);
}
