package tfs
import (
	"../common"
	"fmt"
	"strings"
	b64 "encoding/base64"
)


var REPORTS = []Report {
	& EmptyReport{},
	& ScrumReport{},
}

type Report interface {
	GetName()        string
	GetDescription() string
	Invoke(a *TfsHelper)
}

type TfsHelper struct {
	Username   string
	Pat        string
	Sprint     string
	Team       string
	Endpoint   string
	token      string
}

// TfsHelper functions
func (x *TfsHelper) authenticate() string {
	return b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", x.Username, x.Pat)));
}

func (x *TfsHelper) GetWorkitem(a string, includeHistory bool) map[string]interface{} {
	var rsp = common.InvokeGet(fmt.Sprintf("%s/_apis/wit/workItems/%s", common.BasePath, a), map[string]string{"Authorization":fmt.Sprintf("Basic %s", x.token), "Content-Type":"application/json", "Accepts":"application/json"});
	if includeHistory {
		var rsp2 = common.InvokeGet(fmt.Sprintf("%s/_apis/wit/workItems/%s/updates?api-version=3.1", common.BasePath, a), map[string]string{"Authorization":fmt.Sprintf("Basic %s", x.token), "Content-Type":"application/json", "Accepts":"application/json"});
		rsp["history"] = common.DictionaryToJsonString(rsp2);
	}
	return rsp;
}

func (x *TfsHelper) GetWorkitems(a string, includeHistory bool) []map[string]interface{} {
	var result []map[string]interface{};
	var rsp = common.InvokePost(fmt.Sprintf("%s/%s/_apis/wit/wiql?api-version=1.0", common.BasePath, x.Team), map[string]string{"query":a}, map[string]string{"Authorization":fmt.Sprintf("Basic %s", x.token), "Content-Type":"application/json", "Accepts":"application/json"});
	if rsp["workItems"] != nil {
		for i := 0; i < len(rsp["workItems"].([]interface{})); i++ {
			var wi = rsp["workItems"].([]interface{})[i];
			result = append(result, x.GetWorkitem(fmt.Sprintf("%v", wi.(map[string]interface{})["id"]), includeHistory));
		}
	}
	return result;
}

func (x *TfsHelper) GetPullRequests() []map[string]interface{} {
	var result  []map[string]interface{};
	var rsp = common.InvokeGet(fmt.Sprintf("%s/_apis/git/pullrequests?api-version=3.1", common.BasePath), map[string]string{"Authorization":fmt.Sprintf("Basic %s", x.token), "Content-Type":"application/json", "Accepts":"application/json"});
	for i := 0; i < len(rsp["value"].([]interface{})); i++ {
		result = append(result, rsp["value"].([]interface{})[i].(map[string]interface{}));
	}
	return result;
}

func (x *TfsHelper) GetPullRequestStories() []string {
	var result []string;
	var prs = x.GetPullRequests();
	for _, pr := range  prs {
		var rsp = common.InvokeGet(fmt.Sprintf("%v/workitems", pr["url"]), map[string]string{"Authorization":fmt.Sprintf("Basic %s", x.token), "Content-Type":"application/json", "Accepts":"application/json"});
		for i := 1; i < len(rsp["value"].([]interface{})); i++ {
			var cursor = rsp["value"].([]interface{})[i].(map[string]interface{});
			result = append(result, fmt.Sprintf("%v", cursor["id"]));
		}
	}
    return result;
}
/********************************************************************************************/

// Report functions (DON'T FORGET TO ADD THE REPORT TO THE STATIC MAP)
// 1. Add the struct
// 2. Implement the interface
type EmptyReport struct{}
func (x *EmptyReport) GetName() string { return "EmptyReport"; }
func (x *EmptyReport) GetDescription() string { return "This is just the default implementation for the interface."; }
func (x *EmptyReport) Invoke(a *TfsHelper) {}
type ScrumReport struct{}
func (x *ScrumReport) GetName() string { return "ScrumReport"; }
func (x *ScrumReport) GetDescription() string { return "Generates a report of the statuses for the sprint workitems."; }
func (x *ScrumReport) Invoke(a *TfsHelper) {
	var query = fmt.Sprintf(`Select [System.Id], [System.AssignedTo] From WorkItems Where [System.State] = 'Active' 
		And [Microsoft.VSTS.Common.Triage] <> 'Pending Verification' And [System.WorkItemType] = 'User Story' 
		And [System.IterationPath] Under '%s' And [Microsoft.VSTS.Common.Triage] Not Contains 'Pending QA' Order by [System.AssignedTo] ASC`, a.Sprint);
	var wis = a.GetWorkitems(query, true);
	var prs = a.GetPullRequestStories();
	for _, story := range wis {
		var triage = strings.Trim(fmt.Sprintf("%v", story["fields"].(map[string]interface{})["Microsoft.VSTS.Common.Triage"]), " ");
		var statusDescription = "";
		if common.StrToDictionary([]byte(fmt.Sprintf("%v", story["fileds"])))["Microsoft.VSTS.Common.Triage"] == "Troubleshooting" && common.StrToDictionary([]byte(fmt.Sprintf("%v", story["fields"])))["Custom.UserStory.SupportStatus"] != "" {
			statusDescription = fmt.Sprintf("%v", common.StrToDictionary([]byte(fmt.Sprintf("%v", story["fields"])))["Custom.UserStory.SupportStatus"]);
		}
		if triage == "Pending Merge" {
			statusDescription = "Merging changes";
		} else if triage == "Approved & Refined" {
			var isPrsStory = false;
			for _, cursor := range prs {
				if cursor == story["id"] { isPrsStory = true; break; }
			}
			if isPrsStory {
				statusDescription = "Put out changes, waiting on approvers";
			} else {
				statusDescription = "Still looking into it";
			}
		}
		println(fmt.Sprintf("%v\t%s\t%v\t%v", story["id"], strings.Split(fmt.Sprintf("%v", story["fields"].(map[string]interface{})["System.AssignedTo"].(map[string]interface{})["displayName"]), " ")[0], triage, statusDescription));
	}
}
/********************************************************************************************/

func (x *TfsHelper) Invoke(a string) {
	x.token = x.authenticate();

	var reportFound = false;
	for _, report := range REPORTS {
		if report.GetName() == a {
			reportFound = true;
			report.Invoke(x);
		}
	}
	if !reportFound {
		println("Report (" + a + ") not yet implemented...");
	}
}

