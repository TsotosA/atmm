package main

import "os"

type Exit struct{ Code int }

// HandleExit exit code handler
//https://stackoverflow.com/questions/27629380/how-to-exit-a-go-program-honoring-deferred-calls
func HandleExit() {
	if e := recover(); e != nil {
		if exit, ok := e.(Exit); ok == true {
			os.Exit(exit.Code)
		}
		panic(e) // not an Exit, bubble up
	}
}
