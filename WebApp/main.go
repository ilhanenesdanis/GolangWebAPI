package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

func (hum Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hum.FirstName = "İlhan"
	hum.LastName = "Daniş"
	hum.Age = 23

	//formu parse etmek için kullanılır
	r.ParseForm()

	//sunucudan form bilgisini almak için
	fmt.Println(r.Form)

	//URL'in path bilgisini almak için
	fmt.Println("path", r.URL.Path[1:])

	fmt.Fprint(w, "<table><tr><th>Ad</th><th>Soyad</th><th>Yaş</th></tr><tr><td>"+hum.FirstName+"</td><td>"+hum.LastName+"</td><td>"+strconv.Itoa(hum.Age)+"</td></tr></table>")
}

func main() {
	var hum Human

	err := http.ListenAndServe("localhost:9000", hum)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
