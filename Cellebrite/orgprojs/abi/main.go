package main
/*
#include <encoder.h>
#cgo linux LDFLAGS: -L"./resources/" -lencoder -ldl
#cgo darwin LDFLAGS: -L"./resources/" -lencoder
#cgo linux CFLAGS: -I"/usr/include/" -std=c11
#cgo darwin CFLAGS: -I"/usr/local/include/" -I"./resources/" -std=c11
*/
import "C"
import (
	"fmt"
	"os"
	"reflect"
	"unsafe"
	"strconv"
)

func CStr(s string) *C.char {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return (*C.char)(unsafe.Pointer(h.Data))
}

func ArgsToDictionary(a []string) map[string]string {
	var result = map[string]string{};
	for i := 0; i < len(a) - 1; i++ {
		result[a[i]] = a[i + 1];
	}
	return result;
}

func main() {
	var params = ArgsToDictionary(os.Args);
	var str        = params["--str"];
	var operation  = params["-o"];
	if operation == "" { operation = "encode"; }
	var alg,_      = strconv.Atoi(params["--alg"]);
	if alg < 1 { alg = 1; }
	var rsp *C.char;
	switch (operation) {
		case "decode":
			rsp = C.getDecoded(CStr(str), C.int(alg));
			break;
		case "encode": 
			rsp = C.getEncoded(CStr(str), C.int(alg));
			break;
		default:
			println(fmt.Sprintf("Invalid operation (%s)", operation));
			panic(20);
	}
	println(fmt.Sprintf("%s", C.GoString(rsp)));
	os.Exit(0);
}

