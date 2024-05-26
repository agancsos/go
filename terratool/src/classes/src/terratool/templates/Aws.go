package templates
import (
	"fmt"
)

func AwsTerraformTemplates(federated bool) map[string][]interface{} {
	var rst = map[string][]interface{} {
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
	};
	return rst;
}

func GenerateAwsTaskTemplate(config map[string]interface{}) []interface{} {
	var rst = []interface{} {};
	if EnsureRequiredFields(config) != nil {
		return rst;
	}
	for i := 0; i < config["count"].(int); i++ {
		var temp = map[string]interface{} {
			"@name": config["@name"],
			"@resourceType": fmt.Sprintf("aws_%s", config["@resourceType"]),
		}
		for k, v := range config {
			if k == "@name" || k == "@resourceType" || k == "count" {
				continue;
			}
			temp[k] = v;
		}
		rst = append(rst, temp);
	}
	return rst;
}

