package builtinlogger

import (
	"fmt"
	"log"
)

func Run() {
	log.Println("This is Print")
	// log.Panicln("This is Panic")
	log.Fatalln("This is Fatal")

	fmt.Println("After")

}
