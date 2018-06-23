package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	jst  *time.Location
	once sync.Once
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339))

	nowUTC := now.UTC()
	fmt.Println(nowUTC.Format(time.RFC3339))

	// This fails with "panic: unknown time zone Asia/Tokyo" at play.golang.org
	// jst, err := time.LoadLocation("Asia/Tokyo")
	//if err != nil {
	//  panic(err)
	//}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	nowJST := nowUTC.In(jst)
	fmt.Println(nowJST.Format(time.RFC3339))

	fmt.Println(time.Date(nowJST.Year(), nowJST.Month(), nowJST.Day(), nowJST.Hour(), nowJST.Minute(), nowJST.Second(), 0, time.UTC))
	now1 := time.Now()
	after := now1.Add(10 * time.Hour)
	fmt.Println("\nAdd 10 Hour:", after)

	fmt.Println("")

	t := time.Date(2018, time.May, 22, 12, 0, 0, 0, JST())
	fmt.Println(t)

	fmt.Println(time.Now())

	fmt.Println(time.Now().After(t))
}

// JST jst
func JST() *time.Location {
	once.Do(func() {
		location, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic("couldn't load location.")
		}
		jst = location
	})
	return jst
}
