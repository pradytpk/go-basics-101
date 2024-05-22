package main

import (
	"csv/demo/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/report", controllers.ReportHandler)
	http.HandleFunc("/report/csv", controllers.ReportCsvHandler)
	http.HandleFunc("/upload/csv", controllers.ReportCsvDownloadHandler)

	port := ":8080"
	log.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("could not start server:%s\n", err)
	}
}
