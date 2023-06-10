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
    var params     = common.ArgsToDictionary(os.Args);
    var config map[string]interface{};
    var configFile = params["-f"];
    var query      = params["--sql"];
    if configFile == "" {
        println("\033[31mMust provide connection string file path...\033[m");
        os.Exit(1);
    }
    if query == "" {
        println("\033[31mMust provide query string...\033[m");
        os.Exit(2);
    }
    var rawConfig, err = ioutil.ReadFile(configFile);
    if err != nil {
        println(fmt.Sprintf("\033[31m%s\033[m", err));
        os.Exit(3);
    }
    err = json.Unmarshal(rawConfig, &config);
    var connection = &data.DataConnectionOdbc{};
    connection.SetUsername(config["username"].(string));
    decoded, err := base64.StdEncoding.DecodeString(config["password"].(string))
    if err != nil {
        println(fmt.Sprintf("\033[31m%s\033[m", err));
        os.Exit(4);
    }
    connection.SetPassword(string(decoded));
    connection.SetConnectionString(config["connectionString"].(string));
    table, err := connection.Query(query);
    if err != nil {
        println(fmt.Sprintf("\033[31m%s\033[m", err));
        os.Exit(5);
    }
    var columns = connection.GetColumnNames(query);
    for i, column := range columns {
        if i > 0 { print(fmt.Sprintf("\033[35m,\033[m")); }
        print(fmt.Sprintf("\033[35m%s\033[m", column));
    }
    println();
    for i, row := range table.Rows() {
        if i > 0 { print(fmt.Sprintf("\033[36m\n\033[m")); }
        for j, column := range columns {
            if j > 0 { print(fmt.Sprintf("\033[36m,\033[m")); }
            print(fmt.Sprintf("\033[36m%s\033[m", row.Column(column).Value()));
        }
    }
    println();
    os.Exit(0);
}
