package common
import "C"
import (
	"unsafe"
	"reflect"
	exec "os/exec"
	"fmt"
	"strings"
	"encoding/json"
	"unicode"
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

func CleanString(a string) string {
	var result = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, a);
	return result;
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

func RunCmdNoWait(cmd string) {
	args := strings.Fields(cmd);
    var cmd2 = exec.Command(args[0], args[1:]...);
	cmd2.Run();
	cmd2.Process.Release();
}

func RunCmd(cmd string) string {
	args := strings.Fields(cmd);
	var cmd2 = exec.Command(args[0], args[1:]...);
	var pipe,_ = cmd2.StdoutPipe();
	var stdout = make([]byte, 1024);
	cmd2.Start();
	pipe.Read(stdout)
	cmd2.Process.Release();
	return strings.Replace(string(stdout), "\x00", "", -1);
}

func ArgsToDictionary(a []string) map[string]string {
	var result = map[string]string{};
	for i := 0; i < len(a) - 1; i++ {
		result[a[i]] = a[i + 1];
	}
	return result;
}

func LookupParameterValue(key string, parameters map[string]string, defaultValue string) string {
    if parameters[key] != "" {
        return parameters[key];
    }
	return defaultValue;
}

func DictionaryValue(dict map[string]interface{}, key string, defaultValue interface{}) interface{} {
    if dict[key] == nil {
        return defaultValue;
    }
    return dict[key];
}

func StrToDictionary(s []byte) map[string]interface{} {
	var obj map[string]interface{};
	json.Unmarshal(s, &obj);
	return obj;
}

func DictionaryToJsonString (a map[string]interface{}) string {
	var result = "{";
	var i = 0;
	for key, value := range a {
		if i > 0 { result += ","; }
		var value2 = strings.Replace(fmt.Sprintf("%v", value), "\"", "\\\"", -1);
		result += fmt.Sprintf("\"%s\":\"%v\"", key, value2);
		i++;
	}
	result += "}";
	return result;
}

func StrDictionaryToJsonString (a map[string]string) string {
    var result = "{";
	var i = 0;
    for key, value := range a {
		if i > 0 { result += ","; }
        result += fmt.Sprintf("\"%s\":\"%s\"", key, value);
		i++;
    }
    result += "}";
    return result;
}

func ToConstStr(a string) *C.uchar {
	return (*C.uchar)(unsafe.Pointer(&[]byte(a)[0]))
}

func StrToStrDictionary(s string) map[string]string {
	var rawDict = StrToDictionary([]byte(s));
	var result = map[string]string {};
	for key, value := range rawDict {
		result[key] = fmt.Sprintf("%v", value);
	}
	return result;
}

func DictionaryToStrDictionary(a map[string]interface{}) map[string]string {
	var dict = map[string]string{};
	for key, value := range a {
		dict[key] = fmt.Sprintf("%v", value);
	}
	return dict;
}

func BoolToInt(a bool) int {
	if a {
		return 1;
	} else {
		return 0;
	}
}

func StrToStrPtr(a string) *string {
	var result *string
	temp := a;
	result = &temp;
	return result;
}

func IntToIntPtr(a int32) *int32 {
	var result *int32;
	temp := a;
	result = &temp;
	return result;
}

func BoolToBoolPtr(a bool) *bool {
    var result *bool;
    temp := a;
    result = &temp;
    return result;
}
