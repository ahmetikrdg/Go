package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	go func() {
		for i := 4; i < 7; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total time for Execution: ", elapsedTime)
	time.Sleep(time.Second)
	//başlarına go koymadan 1.2683ms'de işlemi bitirdi.
	//ilk önce üstteki koşuyo takibende alttaki koşuyor.
	//func'un başına go yazdım. Çalıştırdığımda takiben şöyle bişey yapar
	//go keywordünü kullandım ve şuanki durumsa 49.5µs yani 49 mikrosaniye
	//bunun üzerinden giderken mevcut iş yükümüzü eş zamanlı olarak paylaştırdık. go olayı tam olarak böyle.

}
