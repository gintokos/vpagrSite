package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/gintokos/vpagrSite/internal/app"
)

func main() {
	go app.MustStartApp()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	app.Stop(true)
	log.Println("App has stopped his work")
}
