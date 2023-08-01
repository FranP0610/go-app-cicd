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
	ipv4Address []string
}

//type Data struct {
//	localIp string
//}

func getEcsMetadata() string {
	response, err := http.Get("http://169.254.170.2/v2/metadata")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		//fmt.Fprintf(os.Stderr, "Error getting ECS Task Metada: %v\n", err)
		//os.Exit(1)
		panic(err)
	}
	var metadata taskMetadata
	if err = json.Unmarshal(body, &metadata); err != nil {
		panic(err)
	}

	if len(metadata.Containers) > 0 && len(metadata.Containers[0].Networks) > 0 && len(metadata.Containers[0].Networks[0].ipv4Address) > 0 {
		return metadata.Containers[0].Networks[0].ipv4Address[0]
	} else {
		return "NoValue"
	}
}

func getIndexHtml(responseWriter http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	ipAddress := getEcsMetadata()
	//if ipAddress == "NoValue" {
	//	ipA = ipAddress
	//}
	//fmt.Println(user)
	if err != nil {
		panic(err)
	} else {
		template.Execute(responseWriter, ipAddress)
	}

}
func main() {
	http.HandleFunc("/app", getIndexHtml)
	// Initialize web server on port 8080 without error handler
	http.ListenAndServe(":8000", nil)

}
