package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/inancgumus/screen"
)

var zero = []string{
	" ### ",
	"#   #",
	"#   #",
	"#   #",
	" ### ",
}
var one = []string{
	"#    ",
	"#    ",
	"#    ",
	"#    ",
	"#    ",
}

var two = []string{
	"#####",
	"    #",
	"#####",
	"#    ",
	"#####",
}

var three = []string{
	"#####",
	"    #",
	"#####",
	"    #",
	"#####",
}

var four = []string{
	"#   #",
	"#   #",
	"#####",
	"    #",
	"    #",
}

var five = []string{
	"#####",
	"#    ",
	"#####",
	"    #",
	"#####",
}

var six = []string{
	"#####",
	"#    ",
	"#####",
	"#   #",
	"#####",
}

var seven = []string{
	"#####",
	"    #",
	"    #",
	"    #",
	"    #",
}

var eight = []string{
	"#####",
	"#   #",
	"#####",
	"#   #",
	"#####",
}

var nine = []string{
	"#####",
	"#   #",
	"#####",
	"    #",
	"#####",
}

var splitterFull = []string{
	"     ",
	"  #  ",
	"     ",
	"  #  ",
	"     ",
}

var splitterEmpty = []string{
	"     ",
	"     ",
	"     ",
	"     ",
	"     ",
}

var DIGITS = [10][]string{zero, one, two, three, four, five, six, seven, eight, nine}

func printDigits(timeClock *TimeClock) {
	screen.Clear()
	screen.MoveTopLeft()
	var splitter []string
	if !timeClock.getSplitter() {
		splitter = splitterEmpty
	} else {
		splitter = splitterFull
	}

	writer := bufio.NewWriterSize(os.Stdout, 512)
	writer.WriteString("\033[31m")
	currTime := timeClock.getCurrentTime()
	for row := range DIGITS[0] {
		for digit_pos := range currTime {
			digit_num := currTime[digit_pos]
			digit_row := DIGITS[digit_num][row]
			writer.WriteString(digit_row + " ")
			if digit_pos == 1 || digit_pos == 3 {
				writer.WriteString(splitter[row] + " ")
			}
		}
		writer.WriteString("\n")
	}
	writer.WriteString("\033[0m\n")
	writer.Flush()
	// fmt.Printf("%v\033[0m\n", my_screen)
}

var currentTime *TimeClock = &TimeClock{time: make([]int, 6)}
var tempArr [6]int

type TimeClock struct {
	time             []int
	splitter_visible bool
	mu               sync.Mutex
}

func (tc *TimeClock) tryUpdateTime(updater chan interface{}) {
	ticker := time.NewTicker(time.Millisecond)
	for range ticker.C {
		t := time.Now()
		second := tc.getTimeElement(5)
		if currSec := t.Second() % 10; currSec != second {
			tempArr[0] = t.Hour() / 10
			tempArr[1] = t.Hour() % 10
			tempArr[2] = t.Minute() / 10
			tempArr[3] = t.Minute() % 10
			tempArr[4] = t.Second() / 10
			tempArr[5] = currSec
			tc.mu.Lock()
			tc.time = tempArr[:]
			tc.mu.Unlock()
			updater <- new(interface{})
		}
	}
}

func (tc *TimeClock) getTimeElement(index int) (el int) {
	tc.mu.Lock()
	el = tc.time[index]
	tc.mu.Unlock()
	return
}

func (tc *TimeClock) getCurrentTime() (array []int) {
	tc.mu.Lock()
	array = append(array, tc.time...)
	tc.mu.Unlock()
	return
}

func (tc *TimeClock) toggleSplitter() {
	tc.mu.Lock()
	tc.splitter_visible = !tc.splitter_visible
	tc.mu.Unlock()
}

func (tc *TimeClock) getSplitter() (splitter bool) {
	tc.mu.Lock()
	splitter = tc.splitter_visible
	tc.mu.Unlock()
	return
}

func (tc *TimeClock) tickSplitter(updater chan interface{}) {
	ticker := time.NewTicker(time.Millisecond * 500)
	var signal interface{}
	for range ticker.C {
		currentTime.toggleSplitter()
		updater <- signal
	}
}

func main() {
	runtime.GC()
	updater := make(chan interface{})
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go currentTime.tickSplitter(updater)
	go currentTime.tryUpdateTime(updater)
	for {
		select {

		case <-updater:
			printDigits(currentTime)
		case <-done:
			fmt.Println("Closing context")
			_, cancel := context.WithCancel(context.Background())
			defer cancel()
			return

		}
	}
}
