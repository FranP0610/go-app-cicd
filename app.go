package main

import (
	"encoding/json"
	"html/template"
	"io"
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
	IPv4Addresses []string
}

//type Data struct {
//	LocalIp string
//}

func getEcsLocalAddress() string {
	resp, err := http.Get("http://169.254.170.2/v2/metadata")
	if err != nil {
		panic(err)
	}
	// Close the response body
	defer resp.Body.Close()
	// Read all the contents of the Reader - Buffer to read data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//log.Printf("%s", body)
	// Assigning a variable where to store the body response
	var metadata taskMetadata
	json.Unmarshal([]byte(body), &metadata)
	if len(metadata.Containers[0].Networks[0].IPv4Addresses[0]) > 0 {
		return metadata.Containers[0].Networks[0].IPv4Addresses[0]
	} else {
		return "No Value"
	}

}

func getIndexHtml(responseWriter http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	//localIp := Data{getEcsLocalAddress()}
	// Rendering the html
	if err != nil {
		panic(err)
	} else {
		template.Execute(responseWriter, getEcsLocalAddress())
	}

}
func main() {
	http.HandleFunc("/app", getIndexHtml)
	// Initialize web server on port 8080 without error handler
	http.ListenAndServe(":8080", nil)

}
