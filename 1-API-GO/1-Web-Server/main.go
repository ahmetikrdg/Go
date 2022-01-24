package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		//responseWriter response birşey yazmak için kullanılıyo. sunucudan istemicye gönderilen respona
		rw.Write([]byte("Merhaba")) //response bişe yyazıp istemciye gönderdim. oda benden byte istiyo stringi byte dönüştürdüm
		//neden byte istiyo? burada response ile istemci tarafına istediğim veriyi gönderebilirim onda sıkıntı yok bir dosyada olabilir sıkıntı yok
		//sınır vb olmadığı için genelleme yapılır ve bunları byte olarak verir. network üzerindende webtende veri trafiği byte olarak yapılır
	})
	http.ListenAndServe(":8080", nil)
}
