package gopattern

import (
	"log"
)

func handlePanic() {
	for i := 0; i < 10; i++ {
		defer func() { println(i) }()
		if i == 9 {
			log.Panic("Panic the world")
		}
	}
}
