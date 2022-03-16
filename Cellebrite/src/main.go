package main
import (
	"os"
	"os/signal"
	"time"
	"syscall"
	"fmt"
	"log"
	"strconv"
	"./classes/helpers"
	"./classes/api"
	"./classes/mft"
)

func intHandler(sigChan chan os.Signal) {
	<-sigChan;
	log.Println("Bomb has been deactivated.  The world is saved...");
	os.Exit(0);
}

func main() {
	var params = helpers.ArgsToDictionary(os.Args);
	var port, err = strconv.Atoi(params["--port"]);
	if err != nil {
		log.Println(fmt.Sprintf("Failed to convert port value.  %v", err));
	}
	if port == 0 { port = 3434; }
	var server = api.NewRestService(port, "Cellebrite Rest Service");
	var tool   = params["-t"];
	if tool == "" { tool = params["--tool"]; }

	go server.StartServer();

	if params["-h"] == "1" || params["--help"] == "1" {
		helpers.HelpMenu();
	} else if tool == "" {
		log.Fatalln("Tool cannot be empty...");
	} else {
		switch (tool) {
			case "mft":
				if params["-p"] == "" && params["--path"] == "" {
					log.Fatalln("Path cannot be empty...");
				}
				var path = params["-p"];
				if path == "" { path = params["--path"]; }
				var entry, err = mft.ParseMFT(path, 1024);
				if err != nil {
					log.Fatalln(fmt.Sprintf("%v", err));
				}
				log.Println(entry.Content());
				break;
			case "bomb":
				sigChan := make(chan os.Signal);
				go intHandler(sigChan);
				go func() {
					log.Println("Creating child process...");
					signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM);
				}();
				for i := 300; i >= 0; i-- {
					log.Println(fmt.Sprintf("Bomb will go off in %d seconds.  Press CTRL+C to disarm...", i));
					time.Sleep(1 * time.Second);
				}
				break;
			default: log.Fatalln(fmt.Sprintf("Invalid tool (%s)", tool));
		}
	}
	os.Exit(0);
}

