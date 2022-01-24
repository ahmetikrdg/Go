package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	fileInfo, err := os.Stat("demo.txt") //dosya bilgisi döndürür. eğer dosya yoksa hata döner err ile
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dosya adı: ", fileInfo.Name())
	fmt.Println("byte: ", fileInfo.Size())
	fmt.Println("permission: ", fileInfo.Mode())
	fmt.Println("Son güncelleme tarihi: ", fileInfo.ModTime())
	fmt.Println("bu path bir dizinmi?: ", fileInfo.IsDir())

}
