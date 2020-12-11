package main

import "dates"
import "fmt"

func main(){
	days := 5
	fmt.Println("예약일은 ", days, "안에 예약되었습니다.")
	fmt.Println(days + dates.DaysInWeek)
}


