package main

import (
	"github.com/hungranger/GO_CM_Scanner/pkg/helper"
	"github.com/hungranger/Go_Resume_ChatBot/pkg/injector"
)

func main() {
	server, err := injector.InjectServer()
	helper.CheckError(err)

	err = server.Run()
	helper.CheckError(err)
}
