package main

import (
	"fmt"
	"github.com/PaulDotSH/go-grab-ip"
	"time"
)

func main() {
	go_idle_info.AwaitIdleTime(time.Second * 5)
	fmt.Println("Test from main")
}
