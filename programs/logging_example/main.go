package main

import (
	"fmt"
	// "logging_example/builtinlogger"
	"logging_example/logruslogger"
	// "logging_example/zaplogger"
)

func main() {
	fmt.Println("Logging Example")

	// builtinlogger.Run()
	// builtinlogger.LogInFile()
	// builtinlogger.RunCustomLogger()

	// logruslogger.Run()
	logruslogger.RunLoggingToFile()

	// zaplogger.Run()
}
