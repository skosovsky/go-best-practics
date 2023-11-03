package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan

	switch sig {
	case syscall.SIGHUP:
		gotSIGHUP()
	case syscall.SIGINT:
		gotSIGINT()
	case syscall.SIGTERM:
		gotSIGTERM()
	case syscall.SIGQUIT:
		gotSIGQUIT()
	}
}

func gotSIGHUP() {
	log.Print("SIGHUP handler")
}

func gotSIGINT() {
	log.Print("SIGINT handler")
}

func gotSIGTERM() {
	log.Print("SIGTERM handler")
}

func gotSIGQUIT() {
	log.Print("SIGQUIT handler")
}
