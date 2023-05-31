package helpers
import (
	"os"
	"strings"
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

