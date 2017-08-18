package go_spinner

import (
	"fmt"
	"time"
)

func spinner() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(time.Millisecond * 100)
		}
	}
}
