package main

import (
	"log"
)

func main() {
	log.Println("Starting golang classifier...")

	//
	req := SearchRequest{
		".*.pdf",
		"/home/romain/Downloads2",
		[]Rule{
			[]Match{
				Match{
					"content",
					"(?i)(Crédit Agricole)",
					true,
				},
			},
			{
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
				"copy",
				"/tmp/gros_chien/chien/",
			},
		},
	}

	ch := make(chan UnanalysedFile)
	ch2 := make(chan AnalysedFile)

	for i := 0; i < 1; i++ {
		go analyse(ch, ch2)
		go processFile(ch2)
	}

	search([]SearchRequest{req, req}, ch)
	close(ch)
	<-ch2
}
