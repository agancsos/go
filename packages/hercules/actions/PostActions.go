package actions
import (
	"hercules"
)

func ExecuteHerculesCommand(client *hercules.HerculesClient, cmd string) (map[string]interface{}, error) {
    return client.RawPostRequest("cgi-bin/tasks/syslog", map[string]string{}, map[string]interface{}{
		"command": cmd,
		"send": "Send",
	});
}

