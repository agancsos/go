package jenkins
import "testing"

func TestGetJobs(t *testing.T) {
	var client      = &JenkinsClient{Username: "adminabel", PAT: "11086223c288eb4321b381740207f2142c"};
	var jobs        = GetJobs(client, "http://jenkins:8080", []string{"workflowjob"}, []string{}, false);
	if len(jobs) < 1 {
		t.Fail();
	}
}

