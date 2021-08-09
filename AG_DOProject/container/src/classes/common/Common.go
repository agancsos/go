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

func DictionaryToJsonString (dictionary map[string]interface{}) string {
	var result = "{";
	var i = 0;
	for key, value := range dictionary {
		if i > 0 { result += ","; }
		result += fmt.Sprintf("\"%s\":\"%s\"", key, value);
		i++;
	}
	result += "}";
	return result;
}

func StrDictionaryToJsonString (dictionary map[string]string) string {
    var result = "{";
    for key, value := range dictionary {
        result += fmt.Sprintf("\"%s\":\"%s\"", key, value);
    }
    result += "}";
    return result;
}

func ToConstStr(str string) *C.uchar {
	return (*C.uchar)(unsafe.Pointer(&[]byte(str)[0]))
}
