package main
import (
	"os"
	"common"
	"fmt"
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

func main() {
	var tests = []interface{} {
		map[string]interface{}{
			"terraform": map[string]interface{} {
				"required_providers": map[string]interface{} {
					"aws": map[string]interface{} {
						"@name": "",
						"source": "hashicorp/aws",
					},
					"azure": map[string]interface{} {	
						"@name": "",
						"source": "hashicorp/azure",
					},
				},
			},
		},
		map[string]interface{} {
			"variable": map[string]interface{} {
				"@name": "basic",
				"default": "test",
			},
		},
		map[string]interface{} {
			"variable": map[string]interface{} {
				"@name": "reference",
				"@reference": true,
				"default": "var.test",
			},
		},
		map[string]interface{} {
			"variable": map[string]interface{} {
				"@name": "list",
				"default": []string{ "test1", "test2", },
			},
		},
	};

	var tests2 = map[string][]interface{} {
        "variables_aws": []interface{} {
            map[string]interface{} {
                "variable": map[string]interface{} {
                    "@name": "aws_account",
                    "type": "string",
                    "default": "",
                },
            },
            map[string]interface{} {
                "variable": map[string]interface{} {
                    "@name": "aws_role_name",
                    "type": "string",
                    "default": "",
                },
            },
            map[string]interface{} {
                "variable": map[string]interface{} {
                    "@name": "aws_region",
                    "type": "string",
                    "default": "us-east-1",
                },
            },
        },
        "main_aws": []interface{} {
            map[string]interface{} {
                "provider": map[string]interface{} {
                    "aws": map[string]interface{} {
                        "@name": "",
                    },
                },
            },
        },
		"federated_aws": []interface{} {
			map[string]interface{} {
				"data": map[string]interface{} {
					"@name": "federated",
					"@resourceType": "http",
				},
			},
		},
    };
	
	println(MapToTerraform(tests, 0, 4));

	for k, v := range tests2 {
		println(fmt.Sprintf("\033[35mGenerating: %s\033[m", k));
		println(MapToTerraform(v, 0, 4));
	}

	os.Exit(0);
}
