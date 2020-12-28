package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	// Logger with levels
	InfoLogger = log.New(file, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// Simple logger without level
	// log.SetOutput(file)
	// log.Println("Hello world")
	// log.Println("error occurred ")

	InfoLogger.Println("Application started...")
	InfoLogger.Println("Something is executing...")

	WarningLogger.Println("Something is wrong you should know about it")
	WarningLogger.Println("It may not harm prog execution but better to fix it")

	ErrorLogger.Println("Someting went wrong")
	ErrorLogger.Println("Program failed")
}
