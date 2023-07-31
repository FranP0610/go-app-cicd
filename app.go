package main

import (
	"html/template"
	"net/http"
)

type Data struct {
	LocalIp string
}

//func getLocalIp() (localIp, error) {
//
//}

func getIndexHtml(responseWriter http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	localIp := Data{"10.10.0.0/32"}
	//fmt.Println(user)
	if err != nil {
		panic(err)
	} else {
		template.Execute(responseWriter, localIp)
	}

}
func main() {
	http.HandleFunc("/", getIndexHtml)
	// Initialize web server on port 8080 without error handler
	http.ListenAndServe(":8080", nil)

}
