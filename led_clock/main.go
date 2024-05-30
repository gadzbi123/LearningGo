package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/inancgumus/screen"
	_ "github.com/inancgumus/screen"
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

var middle = []string{
	"     ",
	"  #  ",
	"     ",
	"  #  ",
	"     ",
}

var DIGITS = [10][]string{zero, one, two, three, four, five, six, seven, eight, nine}

func printDigits(timeClock *TimeClock) {
	screen.Clear()
	screen.MoveTopLeft()
	var splitter []string = make([]string, 5)
	copy(splitter, middle)
	if !timeClock.getSplitter() {
		for i, v := range splitter {
			splitter[i] = strings.Replace(v, "#", " ", -1)
		}
	}

	my_screen := "\033[31m"
	currTime := timeClock.getCurrentTime()
	for row := range DIGITS[0] {
		for digit_pos := range currTime {
			digit_num := currTime[digit_pos]
			digit_row := DIGITS[digit_num][row]
			my_screen += fmt.Sprintf("%v ", digit_row)
			if digit_pos == 1 || digit_pos == 3 {
				my_screen += fmt.Sprintf("%v ", splitter[row])
			}
		}
		my_screen += "\n"
	}
	fmt.Printf("%v\033[0m\n", my_screen)
}

var currentTime *TimeClock = &TimeClock{time: make([]int, 6)}

type TimeClock struct {
	time             []int
	splitter_visible bool
	mu               sync.Mutex
}

func (tc *TimeClock) tryUpdateTime(updater chan interface{}) {
	ticker := time.NewTicker(time.Millisecond)
	for range ticker.C {
		var timeArr []int
		t := time.Now()
		second := tc.getTimeElement(5)
		if currSec := t.Second() % 10; currSec != second {
			timeArr = make([]int, 6)
			timeArr[0] = t.Hour() / 10
			timeArr[1] = t.Hour() % 10
			timeArr[2] = t.Minute() / 10
			timeArr[3] = t.Minute() % 10
			timeArr[4] = t.Second() / 10
			timeArr[5] = currSec
			tc.mu.Lock()
			tc.time = timeArr[:]
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
	updater := make(chan interface{})
	go currentTime.tickSplitter(updater)
	go currentTime.tryUpdateTime(updater)
	for {
		<-updater
		printDigits(currentTime)
	}
}
