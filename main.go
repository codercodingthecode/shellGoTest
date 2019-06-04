package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

// const (
// startMsg   =
// spaceMsg   =
// escapeMsg  =
// optionMsg  =
// warningMsg =
// )

type appSettings struct {
	runningTime time.Duration
	tracker     bool
}

type messages struct {
	startMsg, spaceMsg, escapeMsg, optionMsg, warningMsg string
}

// time tick
func tickTock(breaker chan bool, wg *sync.WaitGroup, tk chan int) {
Loop:
	for {
		select {
		case <-breaker:
			wg.Done()
			break Loop

		default:
			if time.Now().Second() == 0 && time.Now().Minute() == 05 {
				tk <- 2 // bong
			} else if time.Now().Second() == 0 {
				tk <- 1 // tock
			} else {
				tk <- 0 // tick
			}
		}
	}

}

// print ticker
func printTk(tk *chan int, options *[3]string) {
	for {
		select {
		case <-*tk:
			fmt.Println(options[<-*tk])
		}
		time.Sleep(1 * time.Second)
	}

}

// kills the application after 3 hours run time
func timeRun(runTime time.Duration) {
	for {
		select {
		case <-time.After(time.Hour * runTime):
			os.Exit(0)
		}
	}
}

func main() {
	// clear up terminal
	screen.Clear()
	screen.MoveTopLeft()

	// setup
	msgs := messages{
		startMsg:   "PRESS ENTER TO START THE CLOCK",
		spaceMsg:   "PRESS SPACE BAR AT ANY TIME TO SELECT AN OPTION DURING RUN TIME",
		escapeMsg:  "PRESS ESC TO EXIT APPLICATION",
		optionMsg:  "CHANGE TICK TO WHATEVER WORD YOU WANT",
		warningMsg: "APPLICATION ALREADY RUNNING, PLEASE CHOOSE ANOTHER OPTION BY PRESSING SPACE OR ESC TO EXIT",
	}

	app := appSettings{
		runningTime: 3,
		tracker:     false,
	}

	var wg sync.WaitGroup
	tk := make(chan int)
	breaker := make(chan bool)
	options := [3]string{"tick", "tock", "bong"}
	// tracker := false

	// setup keyboard key watcher
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	// starting application with user messages
	fmt.Println(msgs.startMsg)
	fmt.Println(msgs.escapeMsg)
	fmt.Println(msgs.spaceMsg)

	wg.Add(1)
	go func() {
		for {
			_, key, err := keyboard.GetKey()

			if err != nil {
				panic(err)
			} else if key == keyboard.KeySpace {
				var str string
				breaker <- true
				fmt.Println(msgs.optionMsg)
				fmt.Scanln(&str)
				fmt.Println("NOW PRINTING ", str, " INSTEAD OF ", options[0])
				options[0] = str
				time.Sleep(time.Second * 3)

				wg.Add(1)
				go tickTock(breaker, &wg, tk)

			} else if key == keyboard.KeyEnter {
				if !app.tracker {
					fmt.Println("STARTING CLOCK")
					wg.Add(1)
					app.tracker = true
					go tickTock(breaker, &wg, tk)
				} else {
					fmt.Println(msgs.warningMsg)
				}

			} else if key == keyboard.KeyEsc {
				fmt.Println("EXITING APPLICATION")
				if app.tracker {
					breaker <- true
				}
				time.Sleep(time.Second * 1)
				os.Exit(0)
			}

		}
	}()
	wg.Add(1)
	go timeRun(app.runningTime)
	wg.Add(1)
	go printTk(&tk, &options)
	wg.Wait()
}
