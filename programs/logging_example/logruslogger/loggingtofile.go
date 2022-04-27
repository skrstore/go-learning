package logruslogger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func RunLoggingToFile() {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetLevel(log.DebugLevel)
	log.Debug("This is Debug")
}
