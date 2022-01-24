package main

import (
	"log"
	"os"
)

//dosya oluşturma
//değişkenkleri burada oluşturucam package scope denir buranın ismine

var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("demo.txt")

	if err != nil {
		log.Fatal(err)
	}
}
