package main

import (
	"html/template"
	"io"
	"log"

	//"io/ioutil"
	"net/http"
)

type taskMetadata struct {
	Containers []Container
}
type Container struct {
	Networks []Network
}

type Network struct {
	ipv4Address []string
}

type Data struct {
	LocalIp string
}

func getEcsMetadata() {
	resp, err := http.Get("http://169.254.170.2/v2/metadata")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("%s", body)

	//bodyString := string(body)
	//log.Println(bodyString)

}

func getIndexHtml(responseWriter http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	//localIp := Data{getEcsMetadata()}
	getEcsMetadata()
	localIp := Data{"10.10.0.0/22"}
	log.Println(localIp)

	if err != nil {
		panic(err)
	} else {
		template.Execute(responseWriter, localIp)
	}

}
func main() {
	http.HandleFunc("/app", getIndexHtml)
	// Initialize web server on port 8080 without error handler
	http.ListenAndServe(":8000", nil)

}
