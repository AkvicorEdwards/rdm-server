package maintenance

import (
	"log"
	"os"
	"os/signal"
	"rdm/dam"
)

func ShutDownListener() {
	down := make(chan os.Signal, 1)
	signal.Notify(down, os.Interrupt, os.Kill)
	<-down
	log.Println("Preparing to close")
	dam.LockAll()
	log.Println("Ready to close")
	os.Exit(0)
}

