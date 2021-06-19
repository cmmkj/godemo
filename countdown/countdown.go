package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld = "Go!"
const countdownStart = 3
const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		//time.Sleep(1 * time.Second)
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	//time.Sleep(1 * time.Second)
	sleeper.Sleep()
	fmt.Fprint(out, finalWorld)
}

func main() {
	//Countdown()
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
