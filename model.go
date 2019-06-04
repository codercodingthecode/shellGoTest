package main

import (
	"sync"
	"time"
)

// AppSettings provide base settings to services
type AppSettings struct {
	runningTime time.Duration
	tracker     bool
}

// AppMessages provide messages to services
type AppMessages struct {
	startMsg, spaceMsg, escapeMsg, optionMsg, warningMsg, defaultMsg, exitMsg, startClockMsg string
}

// AppMain provide services to main application
type AppMain struct {
	settings AppSettings
	messages AppMessages
	tk       chan int
	breaker  chan bool
	wg       *sync.WaitGroup
	options  [3]string
}
