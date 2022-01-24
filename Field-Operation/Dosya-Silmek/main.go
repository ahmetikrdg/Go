package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("demo.txt") //dosya bilgisi döndürür. eğer dosya yoksa hata döner err ile
	if err != nil {
		log.Fatal(err)
	}

}
