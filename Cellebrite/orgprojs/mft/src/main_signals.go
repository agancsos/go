package main
import (
	"os"
	"os/signal"
	"time"
	"syscall"
)

func intHandler(sigChan chan os.Signal) {
	<-sigChan;
	println("Entered handler...");
	os.Exit(0);
}

func main() {
	sigChannel := make(chan os.Signal);
	go intHandler(sigChannel);
	go func() {
		println("Creating child process...");
		signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM);
	}();
	time.Sleep(300 * time.Second);
}

