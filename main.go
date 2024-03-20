package main

import (
	"bufio"
	"log"
	"os"

	"github.com/rudrodip/dummylsp/rpc"
)

func main() {
	logger := getLogger("dummylsp.log")
	logger.Println("Starting dummylsp")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Error decoding message: %v", err)
			continue
		}
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Receive message with method %s", method)
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 06666)
	if err != nil {
		panic(err)
	}

	return log.New(logFile, "[dummylsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
