package main

import (
	"log"
)

func main() {
	log.Println("Starting golang classifier...")

	//
	req := SearchRequest{
		"*.pdf",
		"/home/romain/Downloads",
		[]Rule{
			[]Match{
				Match{
					"content",
					"Crédit Agricole",
					true},
			},
		},
	}

	ch := make(chan UnanalysedFile)
	ch2 := make(chan AnalysedFile)

	go analyse(ch, ch2)
	go processFile(ch2)
	search([]SearchRequest{req, req}, ch)
	close(ch)
}
