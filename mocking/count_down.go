package mocking

import (
	"fmt"
	"io"

	// "os" // main func
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer,s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}

	fmt.Fprint(out, finalWord)
}


// func main(){  // main package
// 	sleeper := &DefaultSleeper{}
// 	Countdown(os.Stdout, sleeper)
// }