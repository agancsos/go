package main
import (
	"fmt"
	"os"
	"common"
	"data"
	"encoding/json"
	"io/ioutil"
	"encoding/base64"
)

func main() {
	var params      = common.ArgsToDictionary(os.Args);
	var config map[string]interface{};
	var configFile  = params["-f"];
	var query       = params["--sql"];
	var format      = params["--format"];
	var outputDelim = ",";
	if format == "" {
		format = "CSV";
	}
	if format != "CSV" && format != "MD" {
		println("\033[31mInvalid output format...\033[m");
		os.Exit(1);
	}
	switch (format) {
		case "CSV":
			outputDelim = ",";
			break;
		case "MD": 
			outputDelim = "|";
			break;
		default: break;
	}
	if configFile == "" {
		println("\033[31mMust provide connection string file path...\033[m");
		os.Exit(2);
	}
	if query == "" {
		println("\033[31mMust provide query string...\033[m");
		os.Exit(3);
	}
	var rawConfig, err = ioutil.ReadFile(configFile);
	if err != nil {
		println(fmt.Sprintf("\033[31m%s\033[m", err));
		os.Exit(4);
	}
	err = json.Unmarshal(rawConfig, &config);
	var connection = &data.DataConnectionOdbc{}; 
	connection.SetUsername(config["username"].(string));
	decoded, err := base64.StdEncoding.DecodeString(config["password"].(string))
	if err != nil {
		println(fmt.Sprintf("\033[31m%s\033[m", err));
		os.Exit(5);
	}
	connection.SetPassword(string(decoded));
	connection.SetConnectionString(config["connectionString"].(string));
	table, err := connection.Query(query);
	if err != nil {
		println(fmt.Sprintf("\033[31m%s\033[m", err));
		os.Exit(6);
	}
	var columns = connection.GetColumnNames(query);
	for i, column := range columns {
		if i > 0 || format == "MD" { 
			print(fmt.Sprintf("\033[35m%s\033[m", outputDelim)); 
		}
		print(fmt.Sprintf("\033[35m%s\033[m", column));
	}
	if format == "MD" {
		print(fmt.Sprintf("\033[35m%s\033[m", outputDelim));
	}
	println();
	if format == "MD" {
		for i := range columns {
			_ = i;
			print(fmt.Sprintf("\033[35m%s--\033[m", outputDelim));
		}
		println(fmt.Sprintf("\033[35m%s\033[m", outputDelim));
	}
	for i, row := range table.Rows() {
		if i > 0 { print(fmt.Sprintf("\033[36m\n\033[m")); }
		for j, column := range columns {
			if j > 0 { print(fmt.Sprintf("\033[36m%s\033[m", outputDelim)); }
			print(fmt.Sprintf("\033[36m%s\033[m", row.Column(column).Value()));
		}
		if format == "MD" {
			print(fmt.Sprintf("\033[36m%s\033[m", row.Column(column).Value()));
		}
	}
	println();
	os.Exit(0);
}

