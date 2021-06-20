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
	"unsafe"
)
type SQLiteConnection struct {
	path           string
	handle         *C.sqlite3
}
func (x *SQLiteConnection) GetConnectionString() string { return x.path; }
func (x *SQLiteConnection) GetUsername() string { return ""; }
func (x *SQLiteConnection) GetPassword() string { return ""; }
func (x *SQLiteConnection) SetUsername(a string) {}
func (x *SQLiteConnection) SetPassword(a string) {}
func (x *SQLiteConnection) connect() bool {
	var handle *C.sqlite3
	if C.sqlite3_open(C.CString(x.path), &handle) == C.SQLITE_OK {
		x.handle = handle;
		return true;
	}
	return false;
}
func (x *SQLiteConnection) disconnect() {
	C.sqlite3_close(x.handle);
}
func (x *SQLiteConnection)Query(a string) DataTable {
	var result = DataTable{};
	if x.connect() {
		var statement *C.sqlite3_stmt;
        C.sqlite3_prepare_v2(x.handle, C.CString(a), C.int(len(a)), &statement, nil);
        var columnCount = int(C.sqlite3_column_count(statement));
		for C.sqlite3_step(statement) == C.SQLITE_ROW {
			var row = DataRow{};
			for i := 0; i < columnCount; i++ {
				var column = DataColumn{};
				column.SetName(C.GoString(C.sqlite3_column_name(statement, C.int(i))));
				n := C.sqlite3_column_bytes(statement, C.int(i))
				p := (*C.char)(unsafe.Pointer(C.sqlite3_column_text(statement, C.int(i))))
				column.SetValue(C.GoStringN(p, n));
			}
			result.AddRow(row);
		}
        C.sqlite3_finalize(statement);
        x.disconnect();
    }
	return result;
}

func (x *SQLiteConnection)GetTableNames() []string {
	var result []string;
	if x.connect() {
		results := x.Query("SELECT name FROM sqlite_master WHERE type IN ('table', 'view') ORDER BY 1 DESC");
		for i := range results.GetRows() {
			result = append(result, results.GetRows()[i].GetColumns()[0].GetValue());
		}
        x.disconnect();
    }
    return result;
}

func (x *SQLiteConnection)GetColumnNames(a string) []string {
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

func (x *SQLiteConnection)RunQuery(a string) bool {
	if x.connect() {
		if C.sqlite3_exec(x.handle, C.CString(a), nil, nil, nil) != C.SQLITE_OK {
			return false;
        }
        x.disconnect();
     }
	return true;
}
