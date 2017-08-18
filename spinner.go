//This is show spinner in terminal.
//Insert into code
//go func gospinner.Spinner()
package gospinner

import (
	"fmt"
	"time"
)

//Spinner() run spinner
func Spinner() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(time.Millisecond * 100)
		}
	}
}
