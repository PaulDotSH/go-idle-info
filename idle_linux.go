//go:build linux
// +build linux

package go_idle_info

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

// Only works with xinput installed, and a pretty dumb way but kinda works, anyone who knows of a better way please contact me
func AwaitIdleTime(duration time.Duration) {
	cmd := exec.Command("xinput", "test-xi2", "--root")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Start()

	//while the user is active
	Len := out.Len()
	time.Sleep(time.Millisecond * refresh)

	dur := time.Nanosecond
	// while true, compare the last output from xinput to the current one, if they have different sizes, there was user activity
	for {
		newLen := out.Len()

		//if there was no activity
		if Len == newLen {
			dur = dur + refresh*time.Millisecond
			//clear the buffer
			if Len > 1024*128 {
				fmt.Println("reset")
				p := make([]byte, Len)
				out.Read(p)

				Len = 0
				newLen = 0
			}
		} else { //set both length variables to the same correct length
			Len = newLen
			dur = time.Nanosecond //reset idle time
		}
		//if the user was idle for long enough
		if dur >= duration {
			break
		}
		//fmt.Println(Len) //for debugging only
		time.Sleep(time.Millisecond * refresh)
	}

	if err != nil {
		log.Fatal(err)
	}
}
