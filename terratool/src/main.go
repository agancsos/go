package main
import (
	"fmt"
	"os"
	"common"
	"encoding/json"
	"io/ioutil"
	"terratool"
	"strings"
)

func helpMenu() {
	println(common.PadLeft("", 80, "#"));
	println(fmt.Sprintf("%s#", common.PadLeft(fmt.Sprintf("# Name       : %s", terratool.ApplicationName), 79, " ")));
	println(fmt.Sprintf("%s#", common.PadLeft(fmt.Sprintf("# Author     : %s", terratool.AuthorName), 79, " ")));
	println(fmt.Sprintf("%s#", common.PadLeft(fmt.Sprintf("# Version    : v. %s", terratool.VersionString), 79, " ")));
	println(fmt.Sprintf("%s#", common.PadLeft(fmt.Sprintf("# Description: %s", terratool.DescriptionString), 79, " ")));
	println(fmt.Sprintf("%s#", common.PadLeft("# Flags      :", 79, " ")));
	for k, v := range terratool.Flags {
		println(fmt.Sprintf("# %s: %s #", common.PadLeft(k, 20, " "), common.PadLeft(v, 54, " ")));
	}
	println(common.PadRight("", 80, "#"));
}

func main() {
	var params        = common.ParseParameters(",");
	var config        = map[string]interface{}{};
	var isHelp        = common.LookupParameter(params, "-h", false).(bool);
	var isVersion     = common.LookupParameter(params, "--version", false).(bool);
	var dryRun        = common.LookupParameter(params, "--dry", false).(bool);
	var operation     = common.LookupParameter(params, "--op", "generate").(string);
	var provider      = common.LookupParameter(params, "--provider", "aws").(string);
	var inputPath     = common.LookupParameter(params, "-i", "").(string);
	if inputPath != "" {
		var raw, err = ioutil.ReadFile(inputPath);
		if err != nil {
			println(fmt.Sprintf("\033[33m%s\033[m", err.Error()));
		} else {
			json.Unmarshal(raw, &config);
		}
	}
	for k, v := range params {
		config[strings.Replace(strings.Replace(k, "--", "", -1), "-", "", -1)] = v;
	}
	if isHelp {
		helpMenu();
	} else if isVersion {
		println(fmt.Sprintf("%s v. %s", terratool.ApplicationName, terratool.VersionString));
	} else {
		var service = terratool.NewTerratoolService(config, provider);
		err         := service.Invoke(operation, dryRun);
		if err != nil {
			println(fmt.Sprintf("\033[33m%s\033[m", err.Error()));
			os.Exit(2);
		}
	}
	os.Exit(0);
}

