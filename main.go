package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"time"
)

const (
	pause = "pause"
	work  = "work"
)

var ch chan string

func init() {
	configPath := flag.String("config", "main.ini", "Config file path")
	must(readConfig(*configPath))
	fmt.Printf("Successfully read config on: %s\n", *configPath)

	ch = make(chan string, 1)
}

func main() {
	ch <- work
	pomodoro()
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func pomodoro() {
	for {
		select {
		case str := <-ch:
			msg := ConfigData.Message[str]
			sendNotif(msg)

			time.Sleep(msg.duration)
			ch <- reverse(str)
		}
	}
}

//use command notify-send to send notification
//create slice of arguments based on message config data
//e.g. if useTimeout is true then append -t timeout dur
func sendNotif(msg *MessageConfig) {
	var args []string

	if msg.UseTimeout {
		args = append(args, []string{"-t", fmt.Sprintf("%d", msg.Timeout)}...)
	}

	if msg.UseIcon {
		args = append(args, []string{"-i", msg.Icon}...)
	}

	args = append(args, []string{msg.Title, msg.Message}...)
	cmd := exec.Command("notify-send", args...)
	cmd.Run()
}

//stupid func that basically returns pause and work in turn
//e.g. if input is pause then output is work
func reverse(s string) string {
	if s == pause {
		return work
	}
	return pause
}
