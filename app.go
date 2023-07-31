package main

import (
	"context"
	metadata "github.com/brunoscheufler/aws-ecs-metadata-go"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	LocalIp string
}

func getEcsMetadat() () {
	// Fetch ECS Task metadata from environment
	meta, err := metadata.Get(context.Background(), &http.Client{})
	if err != nil {
		panic(err)
	}
	// Based on the Fargate platform version, we'll have access
	// to v3 or v4 of the ECS Metadata format
	switch m := meta.(type) {
	case *metadata.TaskMetadataV3:
		log.Printf("%s %s:%s", m.Cluster, m.Family, m.Revision)
	case *metadata.TaskMetadataV4:
		log.Printf("%s(%s) %s:%s", m.Cluster, m.AvailabilityZone, m.Family, m.Revision)
	}
}

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
	http.HandleFunc("/app", getIndexHtml)
	// Initialize web server on port 8080 without error handler
	http.ListenAndServe(":8000", nil)
	//call get metadata
	getEcsMetadat()

}
