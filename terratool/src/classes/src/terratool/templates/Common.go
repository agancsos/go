package templates
import (
	"errors"
	"fmt"
)

var RequiredFields = []string{
    "service",
    "type",
    "count",
};

func CommonTerraformTemplates(federated bool) map[string][]interface{} {
	var rst = map[string][]interface{}{
		"providers": []interface{} {
			map[string]interface{} {
				"required_providers": map[string]interface{} {
					"aws": map[string]interface{} {
						"@name": "",
						"source": "hashicorp/asw",
						"version": "5.12.0",
					},
					"http": map[string]interface{} {
						"@name": "",
						"source": "hashicorp/http",
						"version": "3.4.0",
					},
					"azure": map[string]interface{} {
						"@name": "",
						"source": "hashicorp/azure",
						"version": "3.3.3",
					},
				},
			},
		},
	};
	return rst;
}

func EnsureRequiredFields(dict map[string]interface{}) error {
	for _, field := range RequiredFields {
		_, exists := dict[field]; if !exists {
			return errors.New(fmt.Sprintf("Field (%s) is missing"));
		}
	}
	return nil;
}

