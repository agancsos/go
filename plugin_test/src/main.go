package main
import (
	"os"
	"fmt"
	"plugin"
	"plugins"
	"strings"
	"reflect"
	"encoding/json"
	"errors"
	"time"
)

func ParseParameters(delim string) map[string]interface{} {
    var rst = map[string]interface{}{};
    if delim == "" {
        delim = ",";
    }
    for i := range os.Args {
        if len(os.Args[i]) > 0 && string(os.Args[i][0]) == "-" && (i == len(os.Args) - 1 || (i < len(os.Args) - 1 && len(os.Args[i + 1]) > 0 && string(os.Args[i + 1][0]) == "-")) {
            rst[os.Args[i]] = true;
        } else if i < len(os.Args) - 1 && strings.Contains(os.Args[i + 1], delim) {
            rst[os.Args[i]] = strings.Split(os.Args[i + 1], delim);
        } else if i < len(os.Args) - 1 {
            rst[os.Args[i]] = os.Args[i + 1];
        }
    }
    return rst;
}

func LookupParameter(args map[string]interface{}, name string, defaultValue interface{}) interface{} {
    if args[name] == nil {
        return defaultValue;
    }
    return args[name];
}

func LookupAction(pluginPath string, actionName string, values map[string]interface{}) (error, plugins.IAction) {
	module, err    := plugin.Open(pluginPath);
    if err != nil {
        return errors.New("Failed to open plugin library..."), nil;
    }
    sym, err       := module.Lookup(actionName);
    if err != nil {
        return errors.New("Failed to lookup symbal..."), nil;
    }
	for k, v := range values {
    	reflect.ValueOf(sym).Elem().FieldByName(k).Set(reflect.ValueOf(v));
	}
    objInstance := sym.(plugins.IAction);
	return nil, objInstance;
}

func main() {
	var args         = ParseParameters(",");
	var pluginPath   = LookupParameter(args, "-p", "./dist/plugins/test1.so").(string);
	var actionName   = LookupParameter(args, "-a", "Test1Action").(string);
	go func() {
		err, objInstance := LookupAction(pluginPath, actionName, map[string]interface{} {
			"Script": "TEST1",
		})
		if err != nil {
			println(fmt.Sprintf("\033[31m%s\033[m", err.Error()));
			os.Exit(1);
		}
		err, rsp         := objInstance.InvokeAction();
		rspJson, err     := json.Marshal(rsp);
		if err != nil {
			println("\033[31mFailed to serialize to JSON...\033[m");
			os.Exit(2);
		}
		println(fmt.Sprintf("\033[36m%s\033[m", rspJson));
	}();

	go func() { 
		err, objInstance := LookupAction(pluginPath, actionName, map[string]interface{} {
        	"Script": "TEST2",
    	})
    	if err != nil {
        	println(fmt.Sprintf("\033[31m%s\033[m", err.Error()));
        	os.Exit(1);
    	}
    	err, rsp         := objInstance.InvokeAction();
    	rspJson, err     := json.Marshal(rsp);
    	if err != nil {
        	println("\033[31mFailed to serialize to JSON...\033[m");
        	os.Exit(3);
    	}
    	println(fmt.Sprintf("\033[36m%s\033[m", rspJson));
	}();

	time.Sleep(4 * time.Second);
	os.Exit(0);
}

