package countdown

import (
	"fmt"
	"io"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}
