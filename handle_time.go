package gopattern

import "time"

func formatTime() {
	println(time.Now().Format("2006-01-02 15:04:05 MST"))
}
