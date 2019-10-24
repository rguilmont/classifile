package main

import (
	"log"
)

func main() {
	log.Println("Starting golang classifier...")

	req := SearchRequest{
		".*.pdf",
		"/home/romain/Downloads2",
		[]Rule{
			[]Match{

				Match{
					"content",
					"(?i)(ocana)",
					true,
				},
				Match{
					"content",
					"(?i)(angie)",
					false,
				},
			},
		},
		[]Action{
			Action{
				"move",
				"/tmp/gros_chien/chien/",
			},
		},
	}

	ch := make(chan UnanalysedFile)
	ch2 := make(chan AnalysedFile)

	go analyse(ch, ch2)
	go processFile(ch2)

	search([]SearchRequest{req, req}, ch)
	close(ch)
	<-ch2
}
