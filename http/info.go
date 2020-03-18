package http

import (
	"log"
	"os"
	"runtime"
)

func printInfo() {
	log.Printf("PID: \t\t\t\t%d", os.Getpid())
	log.Printf("System: \t\t\t%s", runtime.GOOS)
	log.Printf("System Arch: \t\t%s", runtime.GOARCH)
	log.Printf("Compile Version: \t%s", runtime.Version())
	log.Printf("HTTP Port: \t%d", port)
}
