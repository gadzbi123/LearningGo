package main

import (
	"fmt"
	"strings"
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
func printDigits(digits []int, splitter_visible bool) {
	arr := make([][]string, 5)
	for i := range arr {
		arr[i] = make([]string, 0)
	}
	var my_screen string
	var spl_mid []string
	if splitter_visible {
		spl_mid = getSplittedDigit(middle)
	} else {
		for i := 0; i < 5; i++ {
			spl_mid = append(spl_mid, strings.Repeat(" ", 5))
		}
	}

	for i := 0; i < 5; i++ {
		for _, digit := range digits {
			splitted := getSplittedDigit(DIGITS[digit])
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
	fmt.Println(my_screen)
}
func TimeArray() []int {
	var timeArr []int = make([]int, 6)
	t := time.Now()
	if timeArr[5] == t.Second()%10 {
		return timeArr
	}
	timeArr[0] = t.Hour() / 10
	timeArr[1] = t.Hour() % 10
	timeArr[2] = t.Minute() / 10
	timeArr[3] = t.Minute() % 10
	timeArr[4] = t.Second() / 10
	timeArr[5] = t.Second() % 10
	return timeArr
}

func checkUpdaters(timeArray []int, updater chan []int) {
	timePreBuffer := TimeArray()
	// fmt.Println("timePreBuffer:", timePreBuffer[5], "timeArr", timeArray[5])
	if timePreBuffer[5] != timeArray[5] {
		updater <- timePreBuffer
	}
}

func updateSplitter(updater chan []int) {
	t := time.NewTicker(time.Millisecond * 500)
	<-t.C
	updater <- currentTime
}

var currentTime []int

func main() {
	updater := make(chan []int)
	var splitter_visible bool
	currentTime = TimeArray()
	go updateSplitter(updater)
	for {
		select {
		case temp := <-updater:
			// screen.Clear()
			screen.MoveTopLeft()
			currentTime = temp
			splitter_visible = !splitter_visible
			printDigits(currentTime, splitter_visible)
		default:
			go checkUpdaters(currentTime, updater)
			// time.Sleep(time.Millisecond)
		}
	}

}
