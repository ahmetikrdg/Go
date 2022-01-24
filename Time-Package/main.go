package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Şuanki Unix Zamanı: %v\n", time.Now().Unix())
	//Unix Time: 1 ocak 1970 den beridir geçen saniye sayısına denilen sayısal veri tipi
	time.Sleep(2 * time.Second)
	fmt.Printf("Şuanki Unix Zamanı: %v\n", time.Now().Unix())
}
