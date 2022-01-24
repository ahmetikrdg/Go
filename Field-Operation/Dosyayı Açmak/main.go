package main

import (
	"log"
	"os"
)

func main() {
	/*
		file, err := os.Open("demo.txt")

		if err != nil {
			log.Fatal(err)
		}
		//üzerine veri yazdık diyelim
		//işimiz bittiyse hemen kapamamız lazım
		file.Close()
	*/

	//ikinci parametre dosya açılışını yapar. O_WRONLY sadece yazma işlemi yapar
	//üçüncü parametre dosya izinlerini belirler
	file, err := os.OpenFile("demo.txt", os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
