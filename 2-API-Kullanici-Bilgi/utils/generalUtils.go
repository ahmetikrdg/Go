package utils

import "fmt"

func CheckError(err error) {
	if err != nil {
		fmt.Println("Sıkıntı var")
	}
}
