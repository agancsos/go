package data
/*
#cgo linux LDFLAGS: -lm
#cgo openbsd LDFLAGS: -lm
#cgo linux,!android CFLAGS: -DHAVE_FDATASYNC=1
#cgo linux,!android CFLAGS: -DHAVE_PREAD=1 -DHAVE_PWRITE=1
#cgo darwin CFLAGS: -DHAVE_FDATASYNC=1
#cgo darwin CFLAGS: -DHAVE_PREAD=1 -DHAVE_PWRITE=1
#cgo windows LDFLAGS: -Wl,-Bstatic -lwinpthread -Wl,-Bdynamic
#cgo !windows CFLAGS: -DHAVE_USLEEP=1
#cgo windows,386 CFLAGS: -D_localtime32=localtime
#include <assert.h>
#include <pthread.h>
#include "sqlite3.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)
type DataConnectionSQLite struct {
	path           string
	handle         *C.sqlite3
}
func (x *DataConnectionSQLite) connect() bool {
	var handle *C.sqlite3
	if C.sqlite3_open(C.CString(x.path), &handle) == C.SQLITE_OK {
		x.handle = handle;
		return true;
	}
	return false;
}

func (x *DataConnectionSQLite) disconnect() {
	C.sqlite3_close(x.handle);
}

func (x *DataConnectionSQLite) Query(a string) *DataTable {
	var result = &DataTable{};
	if x.connect() {
		var statement *C.sqlite3_stmt;
        C.sqlite3_prepare_v2(x.handle, C.CString(a), C.int(len(a)), &statement, nil);
        var columnCount = int(C.sqlite3_column_count(statement));
		for C.sqlite3_step(statement) == C.SQLITE_ROW {
			var row = &DataRow{};
			for i := 0; i < columnCount; i++ {
				var column = &DataColumn{};
				column.SetName(fmt.Sprintf("%v", C.GoString(C.sqlite3_column_name(statement, C.int(i)))));
				p := (*C.char)(unsafe.Pointer(C.sqlite3_column_text(statement, C.int(i))))
				column.SetValue(fmt.Sprintf("%v", C.GoString(p)));
				row.AddColumn(column);
			}
			result.AddRow(row);
		}
        C.sqlite3_finalize(statement);
        x.disconnect();
    }
	return result;
}

func (x DataConnectionSQLite) GetTableNames() []string {
	var result []string;
	if x.connect() {
		results := x.Query("SELECT name FROM sqlite_master WHERE type IN ('table', 'view') ORDER BY 1 DESC");
		for i := range results.Rows() {
			result = append(result, results.Rows()[i].Columns()[0].Value());
		}
        x.disconnect();
    }
    return result;
}

func (x DataConnectionSQLite) GetColumnNames(a string) []string {
	var result []string;
	if x.connect() {
		var statement *C.sqlite3_stmt;
        C.sqlite3_prepare_v2(x.handle, C.CString(a), C.int(len(a)), &statement, nil);
		var columnCount = int(C.sqlite3_column_count(statement));
        for i := 0; i < columnCount; i++ {
            result = append(result, C.GoString(C.sqlite3_column_name(statement, C.int(i))));
        }
        C.sqlite3_finalize(statement);
        x.disconnect();
    }
	return result;
}

func (x *DataConnectionSQLite) RunQuery(a string) bool {
	if x.connect() {
		if C.sqlite3_exec(x.handle, C.CString(a), nil, nil, nil) != C.SQLITE_OK {
			return false;
        }
        x.disconnect();
     }
	return true;
}
func (x DataConnectionSQLite) ConnectionString() string { return x.path; }
func (x DataConnectionSQLite) Username() string { return ""; }
func (x DataConnectionSQLite) Password() string { return ""; }
func (x *DataConnectionSQLite) SetUsername(a string) {}
func (x *DataConnectionSQLite) SetPassword(a string) {}
func (x *DataConnectionSQLite) SetConnectionString(a string) { x.path = a; }
