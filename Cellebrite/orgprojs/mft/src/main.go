package main
import (
	"os"
	"fmt"
	"log"
	"./classes/mft"
)

func main() {
	var params = map[string]string{};
	for i := 0; i < len(os.Args) - 1; i++ { params[os.Args[i]] = os.Args[i + 1]; }
	var path = params["-p"];
	var operation = params["-o"];
	if operation == "" { operation = "scan"; };
	switch (operation) {
		case "scan":
			var entry, err = mft.ParseMFT(path, 1024);
			if err != nil {
				log.Fatalln(fmt.Sprintf("%v", err));
			}
			println(entry.Content());
			break;
		default: log.Fatalln(fmt.Sprintf("Invalid operation (%s)", operation));
	}
	os.Exit(0);
}
