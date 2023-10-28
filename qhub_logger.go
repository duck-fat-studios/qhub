package main

import (
	"fmt"

	"github.com/ericboxer/go-kindling"
)


var (
	Logger *kindling.Logger = kindling.NewLogger(kindling.DEBUG)
	consoleLoggerEndpoint *kindling.Endpoint = kindling.NewEndpoint(kindling.DEBUG,false, consoleLoggingFunction)
)

func setupLogger() {

	Logger.RegisterEndpoint(consoleLoggerEndpoint)

}


func consoleLoggingFunction(message string) {
	fmt.Println(message)
}