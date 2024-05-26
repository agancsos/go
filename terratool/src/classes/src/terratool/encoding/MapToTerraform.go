package encoding
import (
	"fmt"
	"common"
	"reflect"
)

func MapToTerraform(items []interface{}, indent int, increment int) string {
    var rst = "";
    for _, item := range items {
        var dict = item.(map[string]interface{});
        for k, v := range dict {
            _, isDict := v.(map[string]interface{}); if isDict {
                rst += common.PadRight("", indent, " ");
				rst += fmt.Sprintf("%s", k);
				_, isResourceTyped := v.(map[string]interface{})["@resourceType"]; if isResourceTyped {
					rst += fmt.Sprintf(" \"%s\" ", v.(map[string]interface{})["@resourceType"].(string));
				}
                _, isNamed := v.(map[string]interface{})["@name"]; if isNamed {
                    if v.(map[string]interface{})["@name"] != "" {
                        rst += fmt.Sprintf(" \"%s\" {\n", v.(map[string]interface{})["@name"].(string));
                    } else {
                        rst += fmt.Sprintf(" = {\n");
                    }
                } else {
                    rst += fmt.Sprintf(" {\n");
                }
                rst += MapToTerraform([]interface{}{ v, }, indent + increment, increment);
                rst += common.PadRight("", indent + increment - increment, " ");
                rst += "}\n";
            } else {
                if k == "@name" || k == "@reference"  || k == "@resourceType"{
                    continue;
                }
                rst += common.PadRight("", indent, " ");
                if reflect.TypeOf(v).Kind() == reflect.Slice {
                    rst += fmt.Sprintf("%s = [", k);
                    if reflect.TypeOf(v).Kind() == reflect.TypeOf([]string{}).Kind() {
                        for i, element := range v.([]string) {
                            if i > 0 {
                                rst += ",";
                            }
                            rst += fmt.Sprintf("\"%v\"", element);
                        }
                    } else if reflect.TypeOf(v).Kind() == reflect.TypeOf([]int{}).Kind() {
                        for i, element := range v.([]int) {
                            if i > 0 {
                                rst += ",";
                            }
                            rst += fmt.Sprintf("%v", element);
                        }
                    }
                    rst += "]";
                } else {
                    if reflect.TypeOf(v).Kind() == reflect.TypeOf(1).Kind() {
                        rst += fmt.Sprintf("%s = %v", k, v);
                    } else if reflect.TypeOf(v).Kind() == reflect.TypeOf("").Kind() {
                        _, isReference := dict["@reference"]; if k == "type" || isReference && dict["@reference"].(bool) {
                        rst += fmt.Sprintf("%s = %v", k, v);
                        } else {
                            rst += fmt.Sprintf("%s = \"%v\"", k, v);
                        }
                    } else {
                        rst += fmt.Sprintf("%s = %v", k, v);
                    }
                }
                rst += "\n";
            }
        }
    }
    return rst;
}

