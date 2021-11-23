package main
import ("fmt"; "os"; "io/ioutil"; "net/http"; "encoding/json")
// Globals
var APPLICATION_NAME	= "disneydownload";
var APPLICATION_AUTHOR  = "Abel Gancsos";
var APPLICATION_VERSION = "1.0.0.0";
var APPLICATION_FLAGS   = map[string]string {
	"-h|--help|-?"	  : "Prints the help menu",
	"-d|--download"	  : "Download video",
	"-i|--info"		  : "Prints info of video",
	"-s|--scan"		  : "Scan from a given key",
};
func PadLeft(str string, le int, pad string) string {
	if len(str) > le { return str[0:le]; }
	result := "";
	for i := len(str); i < le; i++ { result += pad; }
	return str + result;
}
func HelpMenu() {
	println(PadLeft("", 80, "#"));
	println(PadLeft(fmt.Sprintf("# Name        : %s", APPLICATION_NAME), 79, " ") + "#");
	println(PadLeft(fmt.Sprintf("# Author      : %s", APPLICATION_AUTHOR), 79, " ") + "#");
	println(PadLeft(fmt.Sprintf("# Version     : v. %s", APPLICATION_VERSION), 79, " ") + "#");
	println(PadLeft(fmt.Sprintf("# Description : "), 79, " ") + "#");
	println(PadLeft(fmt.Sprintf("# Flags	   :"), 79, " ") + "#");
	for k, v := range APPLICATION_FLAGS { println(PadLeft(fmt.Sprintf("#   %s : %s", k, v), 79, " ") + "#"); }
	println(PadLeft("", 80, "#"));
}
func StrToDictionary(s []byte) map[string]interface{} {
	var obj map[string]interface{};
	json.Unmarshal(s, &obj);
	return obj;
}
func InvokeGet(endpoint string, headers map[string]string) map[string]interface{} {
	var client = http.Client{};
	req, err := http.NewRequest("GET", endpoint, nil);
	for key, value := range headers { req.Header.Add(key, value); }
	rsp, err := client.Do(req);
	if err == nil { rspData, _ := ioutil.ReadAll(rsp.Body); return StrToDictionary(rspData); }
	return nil;
}
/*****************************************************************************/

func main() {
	var operation = "";
	var key	      = "";
	var isHelp	  = false;
	for i := 0; i < len(os.Args); i++ {
		switch (os.Args[i]) {
			case "-h", "--help", "-?": isHelp = true; break;
			case "-d", "--download": operation = "download"; break;
			case "-i", "--info": operation = "info"; break;
			case "-s", "--scan": operation = "scan"; break;
			case "--id": key = os.Args[i + 1]; break;
			default: break;
		}
	}
	if isHelp {
		HelpMenu();
	} else {
		if operation == "" { println("Operation cannot be empty..."); os.Exit(-1); }
		if key == "" { println("Key cannot be empty..."); os.Exit(-2); }
		// https://www.disneyplus.com/video/0123f44d-ccb0-456a-a072-d907cf54574a
		var endpoint = fmt.Sprintf("https://www.disneyplus.com/video/%s", key);
		var rsp = map[string]interface{}{};
		_ = endpoint;
		switch (operation) {
			case "download": break;
			case "info": break;
			case "scan": break;
			default:
				println(fmt.Sprintf("Operation (%s) not supported at this time...", operation));
				os.Exit(-3);
				break;
		}
		println(fmt.Sprintf("%v", rsp));
	}
	os.Exit(0);
}
