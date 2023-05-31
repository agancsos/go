package main
import (
	"os"
	"./classes/helpers"
	"./classes/logsearch/api"
)

func main() {
	var servicePort = helpers.LookupParameter(helpers.ParseParameters(","), "--port", 8080).(int);
	var restService = api.NewRestServer(servicePort);
	restService.StartServer();
	os.Exit(0);
}

