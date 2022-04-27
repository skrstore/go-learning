package builtinlogger

import (
	"log"
	"os"
)

func LogInFile() {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("This is Info")
}
