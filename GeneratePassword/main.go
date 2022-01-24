package main

import (
	"fmt"
	"math/rand"
	"time"
)

//rastgele şifre üretme yada youtube videolarında 6,7 karakterli hash kodu var her birinde farklı o için mesela vb.

var source = rand.NewSource(time.Now().UnixNano())                             //rasgele veri üretmemizi sağlayan nesnem
const _charset = "wertyuopasdfghjklizxcvbnmWERTYUIOPASDFGHKLZXCVBNM0123456789" //bunu kullanarak random değer üreteceğiz

func GeneratePassword(lenght int) string {
	x := make([]byte, lenght)
	for i := range x {
		x[i] = _charset[source.Int63()%int64(len(_charset))] //rastgele değer seçmemizi sağlıyor
	}
	return string(x)
}

func main() {
	password := GeneratePassword(15)
	fmt.Println(password)
}
