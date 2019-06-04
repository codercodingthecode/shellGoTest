package main

import (
	"fmt"
	"os"
	"time"
)

// PrintTk prints ticker
func PrintTk(tk *chan int, options *[3]string) {
	for {
		select {
		case <-*tk:
			fmt.Println(options[<-*tk])
		}
		time.Sleep(1 * time.Second)
	}

}

// TimeRun kills the application after 3 hours run time
func TimeRun(runTime time.Duration) {
	for {
		select {
		case <-time.After(time.Hour * runTime):
			os.Exit(0)
		}
	}
}

// StartupMessages provides guidance to users
func StartupMessages(msgs *AppMessages) {
	fmt.Println(msgs.startMsg)
	fmt.Println(msgs.escapeMsg)
	fmt.Println(msgs.spaceMsg)
}
