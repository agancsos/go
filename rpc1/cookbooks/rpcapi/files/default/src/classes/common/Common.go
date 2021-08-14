package common
import "C"
import (
	"unsafe"
	"reflect"
	exec "os/exec"
	"fmt"
	"strings"
	"encoding/json"
)
func PadRight(str string, le int, pad string) string {
	if len(str) > le {
		return str[0:le];
	}
	result := "";
	for i := len(str); i < le; i++ {
		result += pad;
	}
	return result + str;
}

func PadLeft(str string, le int, pad string) string {
    if len(str) > le {
        return str[0:le];
    }
    result := "";
    for i := len(str); i < le; i++ {
        result += pad;
    }
    return str + result;
}

func CStr(s string) *C.char {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return (*C.char)(unsafe.Pointer(h.Data))
}

func RunCmd(cmd string) string {
	args := strings.Fields(cmd);
	stdout, err := exec.Command(args[0], args[1:]...).Output();
	if err != nil {
		return fmt.Sprintf("%v", err);
	}
	return string(stdout);
}

func StrToDictionary(s []byte) map[string]interface{} {
	var obj map[string]interface{};
	json.Unmarshal(s, &obj);
	return obj;
}

func DictionaryToJsonString (a map[string]interface{}) string {
	var result = "{";
	for key, value := range a {
		result += fmt.Sprintf("\"%s\":\"%v\"", key, value);
	}
	result += "}";
	return result;
}

func StrDictionaryToJsonString (a map[string]string) string {
    var result = "{";
    for key, value := range a {
        result += fmt.Sprintf("\"%s\":\"%s\"", key, value);
    }
    result += "}";
    return result;
}

func ToConstStr(a string) *C.uchar {
	return (*C.uchar)(unsafe.Pointer(&[]byte(a)[0]))
}

func StrToStrPtr(a string) *string {
	var result *string
	temp := a;
	result = &temp;
	temp = "";
	return result;
}

func IntToIntPtr(a int32) *int32 {
	var result *int32;
	temp := a;
	result = &temp;
	return result;
}
