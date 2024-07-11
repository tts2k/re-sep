package utils

import (
	"log"
	"os"

	"re-sep-scraper/config"
)

func InitLogger() {
	log.SetOutput(os.Stdout)
}

func Debug(v ...any) {
	if config.Verbose {
		log.Print(v...)
	}
}

func Debugf(format string, v ...any) {
	if config.Verbose {
		log.Printf(format, v...)
	}
}

func Debugln(v ...any) {
	if config.Verbose {
		log.Println(v...)
	}
}
