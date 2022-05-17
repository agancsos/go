package jenkins
import (
	"strings"
	"fmt"
	"strconv"
)

func contains(list []string, needle string) bool {
	for _, cursor := range list {
		if cursor == needle { return true; }
	}
	return false;
}

func GetJobs(client *JenkinsClient, path string, includeList []string, excludeList []string, recursive bool) ([]map[string]interface{}) {
	var result      = []map[string]interface{}{};
	var jobs, err   = client.JenkinsRequest(path);
	if err == nil {
		var jobs2 = jobs["jobs"].([]interface{});
		for _, job := range jobs2 {
			var jobj   = job.(map[string]interface{});
			if strings.Contains(jobj["_class"].(string), ".folder.") {
				if recursive {
					var temp = GetJobs(client, jobj["url"].(string), includeList, excludeList, recursive);
					for _, t := range temp { result = append(result, t); }
				}
				if contains(includeList, "folder") && !contains(excludeList, "folder") { result = append(result, jobj); }
			} else {
				var comps     = strings.Split(jobj["_class"].(string), ".");
				var className = strings.ToLower(comps[len(comps) - 1]);
				if contains(includeList, className) && !contains(excludeList, className) { result = append(result, jobj); }
			}
		}
	}
	return result;
}

func GetLastBuild(client *JenkinsClient, job map[string]interface{}) (map[string]interface{}, error) {
	return client.JenkinsRequest(fmt.Sprintf("%s/lastBuild", job["url"].(string)));
}

func FindActions(actions map[string]interface{}, className string) interface{} {
	for k, v := range actions {
		if strings.Contains(k, className) { return v; }
	}
	return nil;
}

func GetLogOutput(client *JenkinsClient, job map[string]interface{}, build int) (string, error) {
	if build < 1 {
		var lastBuild, err = GetLastBuild(client, job);
		if err != nil { return "", err; }
		build, _ = strconv.Atoi(fmt.Sprintf("%v", lastBuild["id"]));
	}
	return client.RawRequest(fmt.Sprintf("%s/%d/consoleText", job["url"].(string), build));
}
