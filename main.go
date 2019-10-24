package main

import (
	"flag"
	"log"

	"github.com/jinzhu/configor"
)

// Configuration is the configuration format
type Configuration struct {
	SearchOperations []SearchRequest
}

// Args are startup arguments
type Args struct {
	Configuration string
}

func main() {
	log.Println("Starting file classifier...")
	args := new(Args)
	c := new(Configuration)

	flag.StringVar(&args.Configuration, "conf", "config.yml", "Configuration file to load")
	flag.Parse()

	err := configor.Load(c, args.Configuration)

	if err != nil || len(c.SearchOperations) == 0 {
		log.Panicf("Impossible to load configuration file. %v", err)
	}

	ch := make(chan UnanalysedFile)
	ch2 := make(chan AnalysedFile)

	go analyse(ch, ch2)
	go processFile(ch2)

	search(c.SearchOperations, ch)
	close(ch)
	<-ch2
}
