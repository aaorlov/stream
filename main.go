package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/aaorlov/stream/app"
)

func main() {
	prg := app.New()
	if err := prg.Run(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(runtime.NumGoroutine()) // TODO remove in prod
}
