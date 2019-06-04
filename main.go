package main

import (
	"sync"

	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

func main() {
	var wgg sync.WaitGroup

	// clear up terminal
	screen.Clear()
	screen.MoveTopLeft()

	// setup
	app := AppSettings{
		runningTime: 3,
		tracker:     false,
	}

	msgs := AppMessages{
		startMsg:      "PRESS ENTER TO START THE CLOCK",
		spaceMsg:      "PRESS SPACE BAR AT ANY TIME TO SELECT AN OPTION DURING RUN TIME",
		escapeMsg:     "PRESS ESC TO EXIT APPLICATION",
		optionMsg:     "CHANGE TICK TO WHATEVER WORD YOU WANT",
		warningMsg:    "APPLICATION ALREADY RUNNING, PLEASE CHOOSE ANOTHER OPTION BY PRESSING SPACE OR ESC TO EXIT",
		defaultMsg:    "CANNOT BE EMPTY, DEFAULT TO TICK",
		exitMsg:       "EXITING APPLICATION",
		startClockMsg: "STARTING CLOCK",
	}

	application := AppMain{
		settings: app,
		messages: msgs,
		tk:       make(chan int),
		breaker:  make(chan bool),
		options: [3]string{
			"tick",
			"tock",
			"bong",
		},
		wg: &wgg,
	}

	// setup keyboard key watcher
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	// starting application with user messages
	StarupMessages(&application.messages)

	// starts app go routines
	application.wg.Add(3)
	go HandleUserInput(&application)
	go TimeRun(app.runningTime)
	go PrintTk(&application.tk, &application.options)
	application.wg.Wait()
}
