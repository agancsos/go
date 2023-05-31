package main
import (
	"bufio"
    "fmt"
    "os"
    "./classes/helpers"
	"./classes/common"
	"./classes/logsearch/parsers"
)

func main() {
    var baseEndpoint = helpers.LookupParameter(helpers.ParseParameters(","), "--end", "http://localhost:8080").(string);
	var operation    = helpers.LookupParameter(helpers.ParseParameters(","), "--op", "query").(string);
	var query        = helpers.LookupParameter(helpers.ParseParameters(","), "--query", "").(string);
	var ingestFile   = helpers.LookupParameter(helpers.ParseParameters(","), "-f", "").(string);
	if baseEndpoint == "" {
		println(fmt.Sprintf("\033[31mMissing required field...\033[m"));
		os.Exit(1);
	}
	switch (operation) {
		case "ingest":
			if ingestFile == "" {
        		println(fmt.Sprintf("\033[31mMissing required field...\033[m"));
        		os.Exit(2);
    		}
			_, exists := os.Stat(ingestFile); if exists != nil {
				println(fmt.Sprintf("\033[31mFile does not exists...\033[m"));
				os.Exit(3);
			}
			var fhandle, err = os.Open(ingestFile);
			if err != nil {
				println(fmt.Sprintf("\033[31mFailed to open file. %s\033[m", err.Error()));
				os.Exit(4);
			}
			defer fhandle.Close();
			var scanner = bufio.NewScanner(fhandle);
			for scanner.Scan() {
				var rsp = common.InvokePost(fmt.Sprintf("%s/add/", baseEndpoint), map[string]string{
					"source": ingestFile,
					"content": scanner.Text(),
				}, map[string]string{});
				println(fmt.Sprintf("\033[36m%s\033[m", common.DictionaryToJsonString(rsp)));
			}
			break;
		case "query":
			if query == "" {
                println(fmt.Sprintf("\033[31mMissing required field...\033[m"));
                os.Exit(5);
            }
			var err, queryComps = (&parsers.FilterParser{}).Parse(query);
			if err != nil {
				println(fmt.Sprintf("\033[31mFailed to parse query. %s\033[m", err.Error()));
                os.Exit(6);
			}
			var _, queryString = (&parsers.UrlParameterParser{}).Unparse(queryComps);
			var rsp         = common.InvokeGet(fmt.Sprintf("%s/query/%s", baseEndpoint, queryString), map[string]string{});
			println(fmt.Sprintf("\033[36m%s\033[m", common.DictionaryToJsonString(rsp)));
			break;
		case "purge":
			var rsp = common.InvokePost(fmt.Sprintf("%s/purge/", baseEndpoint), map[string]string{}, map[string]string{});
			println(fmt.Sprintf("\033[36m%s\033[m", common.DictionaryToJsonString(rsp)));
			break;
		default:
			break;
	}
	println(fmt.Sprintf("\033[32mDone!\033[m"));
    os.Exit(0);
}
