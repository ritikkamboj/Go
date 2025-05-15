package main

import (
	"fmt"
	"time"
)

func main() {

	currTime := time.Now()
	fmt.Println("curr time is ", currTime)

	fmt.Printf("type of time is %T\n", currTime)

	formatTime := currTime.Format("dd-mm-yyyy") // this way nor works

	fmt.Println("new formated time is ", formatTime)

	formatTime1 := currTime.Format("02-01-2006")
	fmt.Println("new formated time is ", formatTime1)

	// similarly if we need day then we have to type

	formatTime2 := currTime.Format("02-01-2006, Monday")
	fmt.Println("new formated time is ", formatTime2)

	// similarly we have to use for to show time

	formatTime3 := currTime.Format("02-01-2006, Monday 15:04:05")
	fmt.Println("new formated time is ", formatTime3)

	formatTime4 := currTime.Format("2006-01-02, Monday 15:04:05")
	fmt.Println("new formated time is ", formatTime4)

	formatTime5 := currTime.Format("2006/01/02, Monday 15:04:05")
	fmt.Println("new formated time is ", formatTime5)

	// when we havve to use AM aand PM format

	formatTime6 := currTime.Format("2006/01/02, Monday 3:04 PM")
	fmt.Println("new formated time is ", formatTime6)

	// here we ahevt oo convert from str to date format
	dateStr := "2013-11-25"
	layout_str12 := "2006-01-02" // datestr and layout must be in same format like - or /
	formated_time, _ := time.Parse(layout_str12, dateStr)
	fmt.Println("ormated time is ", formated_time)

	// add one more day in curr time

	new_date := currTime.Add(24 * time.Hour)
	fmt.Println("new_date value is ", new_date)

	format_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("formated date is ", format_new_date)

}
