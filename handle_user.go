package main

import (
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

// HandleUserInput handle user interaction and some of the app's logic
func HandleUserInput(app *AppMain) {
	for {
		_, key, err := keyboard.GetKey()

		if err != nil {
			panic(err)
		} else if key == keyboard.KeySpace {
			var str string
			app.breaker <- true
			fmt.Println(app.messages.optionMsg)
			fmt.Scanln(&str)
			fmt.Println("NOW PRINTING ", str, " INSTEAD OF ", app.options[0])
			app.options[0] = str
			time.Sleep(time.Second * 3)
			app.wg.Add(1)
			go TickTock(app.breaker, app.wg, app.tk)

		} else if key == keyboard.KeyEnter {
			if !app.settings.tracker {
				fmt.Println("STARTING CLOCK")
				app.wg.Add(1)
				app.settings.tracker = true
				go TickTock(app.breaker, app.wg, app.tk)
			} else {
				fmt.Println(app.messages.warningMsg)
			}

		} else if key == keyboard.KeyEsc {
			fmt.Println("EXITING APPLICATION")
			if app.settings.tracker {
				app.breaker <- true
			}
			time.Sleep(time.Second * 1)
			os.Exit(0)
		}
	}
}