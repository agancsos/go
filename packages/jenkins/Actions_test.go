package jenkins
import "testing"

func TestGetJobs(t *testing.T) {
	var client      = &JenkinsClient{Username: "", PAT: ""};
	var jobs        = GetJobs(client, "http://jenkins:8080", []string{"workflowjob"}, []string{}, false);
	if len(jobs) < 1 {
		t.Fail();
	}
}

