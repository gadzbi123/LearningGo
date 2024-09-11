package main

import (
	"fmt"
	"time"
)

func PrintTime(label string, t *time.Time) {
	// fmt.Printf("%s: day: %v, month: %v, year: %v \n", label, t.Day(), t.Month(), t.Year())
	fmt.Println(label, t.Format(time.RFC822))
}
func main() {
	now := time.Now()
	PrintTime("now", &now)
	specific := time.Date(2000, 4, 7, 12, 5, 0, 0, time.Local)
	PrintTime("birth", &specific)
	unix := time.Unix(3252363472, 0)
	PrintTime("random", &unix)
	fmt.Println(unix.Format(time.Kitchen))
	fmt.Println(unix.Format(time.RFC822))
	warsawTime, warError := time.LoadLocation("Europe/Warsaw")
	isError := false
	if warError != nil {
		fmt.Println(warError.Error())
		isError = true
	}
	AngelesTime, AngelesError := time.LoadLocation("America/Los_Angeles")
	if AngelesError != nil {
		fmt.Println(AngelesError.Error())
		isError = true
	}
	if isError {
		return
	}
	fmt.Printf("%+v\n", AngelesTime)
	warsTime, _ := time.ParseInLocation(time.RFC822, time.Now().Format(time.RFC822), warsawTime)
	anglTime, _ := time.ParseInLocation(time.RFC822, time.Now().Format(time.RFC822), AngelesTime)
	PrintTime("Warsaw", &warsTime)
	PrintTime("LosAngles", &anglTime)
}
