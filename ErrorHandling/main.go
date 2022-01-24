package main

import (
	"errors"
	"fmt"
	"os"
)

//Hata Yönetimi
func main() {

	file, err := os.Open("dosyam.txt") //file gönderir hata alırsa err'e gönderir bu standrta alışmalıyız go'da
	//dosyayı okuma sırasında bi hata meydana geldi senaryo
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(file.Name())

	myError := errors.New("Bu bir hata")
	fmt.Println(myError)
}
