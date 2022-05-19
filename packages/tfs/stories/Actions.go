package stories
import (
	"tfs"
	"fmt"
)

func GetWorkItem(client *tfs.TfsClient, path string, wid string, includeHistory bool) map[string]interface{} {
	var workitem, err         = client.TfsRequest(fmt.Sprintf("%s/_apis/wit/workItems/%s", path, wid), map[string]string {"Content-Type":"application/json", "Accepts":"application/json",});
	if err != nil { return nil; }
	if !includeHistory { return workitem; }
	rsp, err                 := client.TfsRequest(fmt.Sprintf("%s/_apis/wit/workItems/%s/updates?api-version=3.1", path, wid), map[string]string {"Content-Type":"application/json", "Accepts":"application/json",});
	if err != nil { return workitem; }
	workitem["history"] = rsp["value"];
	rsp, err                 = client.TfsRequest(fmt.Sprintf("%s/_apis/wit/workItems/%s/history?api-version=3.1", path, wid), map[string]string {"Content-Type":"application/json", "Accepts":"application/json",});
	workitem["comments"] = rsp["value"];
	return workitem;
}

func GetWorkItems(client *tfs.TfsClient, path string, query string, includeHistory bool) []map[string]interface{} {
	var result = []map[string]interface{}{}
	var rsp, err   = client.TfsPostRequest(fmt.Sprintf("%s/%s/_apis/wit/wiql?api-version=1.0", path, client.Team), map[string]string {"Content-Type":"application/json", "Accepts":"application/json",}, map[string]string{"query":query});
	if err != nil { return result; }
	if rsp["workItems"] != nil {
		for _, workitem := range rsp["workItems"].([]interface{}) {
			result = append(result, GetWorkItem(client, path, fmt.Sprintf("%v", workitem.(map[string]interface{})["id"]), includeHistory));
		}
	}
	return result;
}

