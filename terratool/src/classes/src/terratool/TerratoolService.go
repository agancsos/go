package terratool
import (
	"errors"
	"fmt"
	"os"
	"encoding/json"
	"terratool/templates"
	"terratool/encoding"
)

type terratoolService struct {
	provider          string
	outputPath        string
	resources         []map[string]interface{}
	federated         bool
	applicationId     string
}

var __tt_instance__ *terratoolService;
func NewTerratoolService(config map[string]interface{}, provider string) *terratoolService {
	if __tt_instance__ == nil {
		__tt_instance__ = &terratoolService{};
		__tt_instance__.federated     = false;
		__tt_instance__.provider      = provider;
		__tt_instance__.applicationId = config["applicationId"].(string);
		__tt_instance__.outputPath    = config["o"].(string);
		__tt_instance__.resources     = []map[string]interface{}{};
		if config["resources"] != nil {
			_, isList := config["resources"].([]map[string]interface{}); if isList {
				__tt_instance__.resources = config["resources"].([]map[string]interface{});
			} else {
				json.Unmarshal([]byte(config["resources"].(string)), &__tt_instance__.resources);
			}
		}
		if config["federated"] != nil {
			__tt_instance__.federated     = config["federated"].(bool);
		}
	}
	return __tt_instance__;
}

func (x *terratoolService) preparePackage(dryRun bool) {
	_, err := os.Stat(x.outputPath); if err != nil {
		if !dryRun {
			println("\033[35mCreating package directory\033[m");
			os.Mkdir(x.outputPath, 0777);
		}
	}	
	for k, v := range templates.CommonTerraformTemplates(x.federated) {
		println(fmt.Sprintf("\033[35m>> Creating: %s\033[m", k));
		if !dryRun {
		} else {
			println(fmt.Sprintf("\033[36m>>> %s\033[m", encoding.MapToTerraform(v, 0, 4)));
		}
	}
	if x.provider == "aws" {
		for k, v := range templates.AwsTerraformTemplates(x.federated) {
			println(fmt.Sprintf("\033[35m>> Creating: %s\033[m", k));
			if !dryRun {
        	} else {
            	println(fmt.Sprintf("\033[36m>>> %s\033[m", encoding.MapToTerraform(v, 0, 4)));
        	}
		}
	} else {
		for k, v := range templates.AzureTerraformTemplates(x.federated) {
            println(fmt.Sprintf("\033[35>> Creating: %s\033[m", k));
            if !dryRun {
            } else {
                println(fmt.Sprintf("\033[36m>>> %s\033[m", encoding.MapToTerraform(v, 0, 4)));
            }
        }
	}	
}

func (x *terratoolService) prepareResources(dryRun bool) {
	for _, resource := range x.resources {
		var tf = []interface{}{};
		if x.provider == "aws" {
			tf = templates.GenerateAwsTaskTemplate(resource);
		} else {
			tf = templates.GenerateAzureTaskTemplate(resource);
		}
		if !dryRun {
        } else {
            println(fmt.Sprintf("\033[36m>>> {0}\033[m", encoding.MapToTerraform(tf, 0, 4)));
        }
	}	
}

func (x *terratoolService) Invoke(operation string, dryRun bool) error {
	if x.provider != "aws" && x.provider != "azure" {
		return errors.New(fmt.Sprintf("Invalid provider (%s)", x.provider));
	}
	if x.applicationId == "" {
		return errors.New("ApplicationId cannot be empty");
	}
	switch (operation) {
		case "generate":
			x.preparePackage(dryRun);
			x.prepareResources(dryRun);
			break;
		case "purge":
			if !dryRun {
				os.RemoveAll(x.outputPath);
			}
			break;
		default:
			return errors.New(fmt.Sprintf("Invalid operation (%s)", operation));
	}
	return nil;
}


