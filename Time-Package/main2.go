package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2018, time.November, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("Çalışıyor:%s\n", t)

	now := time.Now()
	fmt.Printf("Mevcut zaman: %s\n", now)

	fmt.Println("**************")
	fmt.Println("Ay: ", now.Month())
	fmt.Println("Gün: ", now.Day())
	fmt.Println("Haftanın Günü: ", now.Weekday())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	fmt.Println(now.Nanosecond())
	c := time.Date(2022, time.November, 10, 20, 0, 0, 0, time.UTC)

	println(t.Equal(c)) //t tarihi ve c tarihi birbirine eşitmi. after beforu falanda var

	a := time.Date(2022, time.January, 18, 20, 0, 0, 0, time.UTC)
	fmt.Println("hey: ", a)
	diff := now.Sub(a) //farkı alır. şuanki fark ile a arasında

	println(diff)
	println(diff)

}
