package main

import (
  "fmt"
  "time"
  "os"
  "github.com/beevik/ntp"
)

func Time() {
	timeNTP, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	} else {
		fmt.Print(timeNTP)
		fmt.Print(" / ")
    fmt.Println(time.Now())
	}
}

func main() {
  Time()
}
