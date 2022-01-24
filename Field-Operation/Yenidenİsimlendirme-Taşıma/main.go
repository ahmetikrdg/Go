package main

import (
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	orginalPath := "demo.txt"
	newPath := "test.txt"                  //taşımak için ./bilmemneresi/text.txt yapmak lazım burayı bu şekilde sadece isim değişirt şuanki haliyle
	err := os.Rename(orginalPath, newPath) //eski path yeni path

	if err != nil {
		log.Fatal(err)
	}
}
