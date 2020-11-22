package main

import (
	"fmt"
	"log"
)

// Info method output INFO messages.
func Info(logMsg string) {
	log.Printf(fmt.Sprintf("INFO: %s", logMsg))
}

// Warn method output WARN messages.
func Warn(logMsg string) {
	log.Printf(fmt.Sprintf("WARN: %s", logMsg))
}

// Error method output ERROR messages.
func Error(logMsg string) {
	log.Printf(fmt.Sprintf("ERROR: %s", logMsg))
}
