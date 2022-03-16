package main;
import (
	"os"
	"strconv"
	"./classes/rewards/api"
)

func main() {
	var servicePort = 4441;

	for i, _ := range os.Args {
		switch (os.Args[i]) {
			case "-p", "--port":
				servicePort, _ = strconv.Atoi(os.Args[i + 1]);
				break;
		}
	}
	api.StartServer(servicePort);
	os.Exit(0);
}
