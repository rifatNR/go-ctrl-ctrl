package main

import (
	"encoding/json"
	"fmt"
	"time"

	hook "github.com/robotn/gohook"
)

type response struct {
	Msg string `json:"msg"`
}

func main() {
	// var isCtrlCtrl = false;
	lastCtrlTime := 0

	// hello, _ := json.Marshal(response{Msg: "Hello"})
	// fmt.Println(string(hello))
	// time.Sleep(2 * time.Second)
	// world, _ := json.Marshal(response{Msg: "World"})
	// fmt.Println(string(world))

	hook.Register(hook.KeyDown, []string{"ctrl"}, func(e hook.Event) {
		now := time.Now()
		miliSec := now.UnixMilli()

		timeDiff := int(miliSec) - lastCtrlTime
		if timeDiff < 350 {
			data, _ := json.Marshal(response{Msg: "CTRL + CTRL"})
			fmt.Println(string(data))
			lastCtrlTime = 0
		} else {
			fmt.Print("-")
			lastCtrlTime = int(miliSec)
		}

		// hook.End()
	})

	s := hook.Start()
	<-hook.Process(s)

	for {
		time.Sleep(2 * time.Second)
	}
}
