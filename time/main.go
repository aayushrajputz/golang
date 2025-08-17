package main

import (
	"fmt"
	"time"
)

func main() {
	currenttime := time.Now()
	fmt.Println(currenttime)
	fmt.Printf("type of curenttime is %T\n", currenttime)

	formatted := currenttime.Format("02-01-2006,Monday, 3:04 PM")
	fmt.Println(formatted) 
   
	 layout_strc := "2006-01-02 "
	 dateStr := "2002-05-06"
	 formatted_time,_ := time.Parse	(layout_strc, dateStr)
	 fmt.Println(formatted_time	)
   
	  newdate := currenttime.Add(48 *time.Hour)
	  fmt.Println(newdate)
	  formatted_new_date := newdate.Format("02-01-2006,Monday, 3:04 PM")
	  fmt.Println(formatted_new_date)
}
