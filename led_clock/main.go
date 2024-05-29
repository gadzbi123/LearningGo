package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/inancgumus/screen"
	_ "github.com/inancgumus/screen"
)

const zero string = `
 ### 
#   #
#   #
#   #
 ### 
`
const one string = `
#    
#    
#    
#    
#    
`
const two string = `
#####
    #
#####
#    
#####
`
const three string = `
#####
    #
#####
    #
#####
`

const four string = `
#   #
#   #
#####
    #
    #
`

const five string = `
#####
#    
#####
    #
#####
`

const six string = `
#####
#    
#####
#   #
#####
`

const seven string = `
#####
    #
    #
    #
    #
`

const eight string = `
#####
#   #
#####
#   #
#####
`

const nine string = `
#####
#   #
#####
    #
#####
`
const middle string = `
     
  #  
     
  #  
     
`

var DIGITS = [10]string{zero, one, two, three, four, five, six, seven, eight, nine}

func AssertEqual(statement bool, format string, args ...interface{}) {
	if !statement {
		message := fmt.Sprintf(format+"\n", args...)
		panic(message)
	}
}
func getSplittedDigit(digit string) []string {
	spl_digit := strings.Split(digit, "\n")
	return spl_digit[1 : len(spl_digit)-1]
}
func printDigits(timeClock *TimeClock) {
	screen.Clear()
	screen.MoveTopLeft()
	arr := make([][]string, 5)
	for i := range arr {
		arr[i] = make([]string, 0, 6)
	}
	var my_screen string = "\033[31m"
	var spl_mid []string
	if timeClock.getSplitter() {
		spl_mid = getSplittedDigit(middle)
	} else {
		for i := 0; i < 5; i++ {
			spl_mid = append(spl_mid, strings.Repeat(" ", 5))
		}
	}

	for i := 0; i < 5; i++ {
		for _, timeDigit := range timeClock.getCurrentTime() {
			splitted := getSplittedDigit(DIGITS[timeDigit])
			arr[i] = append(arr[i], splitted[i])
		}

		for j := 0; j < 6; j++ {
			digit_row := arr[i][j]
			my_screen += fmt.Sprintf("%v ", digit_row)
			AssertEqual(len(digit_row) == 5, "Lenght of row should be 5")
			if j == 1 || j == 3 {
				my_screen += fmt.Sprintf("%v ", spl_mid[i])
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

func (tc *TimeClock) updateTime(updater chan interface{}) {
	var timeArr []int
	t := time.Now()
	second := tc.getTimeElement(5)
	if second != t.Second()%10 {
		timeArr = make([]int, 6)
		timeArr[0] = t.Hour() / 10
		timeArr[1] = t.Hour() % 10
		timeArr[2] = t.Minute() / 10
		timeArr[3] = t.Minute() % 10
		timeArr[4] = t.Second() / 10
		timeArr[5] = t.Second() % 10
		tc.mu.Lock()
		tc.time = timeArr[:]
		tc.mu.Unlock()
		updater <- new(interface{})
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

func updateSplitter(updater chan interface{}) {
	ticker := time.NewTicker(time.Millisecond * 500)
	var signal interface{}
	for range ticker.C {
		currentTime.toggleSplitter()
		updater <- signal
	}
}

func main() {
	updater := make(chan interface{})
	go updateSplitter(updater)
	for {
		select {
		case <-updater:
			printDigits(currentTime)
		default:
			go currentTime.updateTime(updater)
			time.Sleep(time.Millisecond)
		}
	}

}
