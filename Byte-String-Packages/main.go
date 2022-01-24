package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	var x bytes.Buffer
	builder := strings.Builder{} //bu daha performanslı bufferdan

	for i := 0; i < 10; i++ {
		builder.WriteString("Data")
	}

	fmt.Println(builder.String())

	for i := 0; i < 10; i++ {
		x.WriteString(rastgeleString()) //aşağıda oluşturduğum rastgele string fonksiyonunu çağırdım gelen veriyi x bufferına string olarak ekle dedim
	}
	fmt.Println(x.String())
}

func rastgeleString() string {
	return "$xyz-3642"
}
