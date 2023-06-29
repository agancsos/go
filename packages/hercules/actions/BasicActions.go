package actions
import (
	"hercules"
)

func GetVersion(client *hercules.HerculesClient) (map[string]interface{}, error) {
	return client.RawRequest("cgi-bin/debug/version_info", map[string]string{});
}

func GetSysLog(client *hercules.HerculesClient, maxLines int) (map[string]interface{}, error) {
    return client.RawPostRequest("cgi-bin/tasks/syslog", map[string]string{}, map[string]interface{}{
		"msgcount": maxLines,
		"send": "Send",
	});
}

func GetCPUConfig(client *hercules.HerculesClient) (map[string]interface{}, error) {
    return client.RawRequest("cgi-bin/configure/cpu", map[string]string{});
}

func GetGeneralRegisters(client *hercules.HerculesClient) (map[string]interface{}, error) {
    return client.RawRequest("cgi-bin/registers/general", map[string]string{});
}

func GetControlRegisters(client *hercules.HerculesClient) (map[string]interface{}, error) {
    return client.RawRequest("cgi-bin/registers/control", map[string]string{});
}

func ListDevices(client *hercules.HerculesClient) (map[string]interface{}, error) {
	return client.RawRequest("cgi-bin/debug/device/list", map[string]string{});
}

