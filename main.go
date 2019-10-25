package main

import (
	"flag"
	"log"

	"github.com/jinzhu/configor"
)

// Global flags
var dryRun = false

// Configuration is the configuration format
type Configuration struct {
	SearchOperations []SearchRequest `required:"true"`
}

// Args are startup arguments
type Args struct {
	Configuration string
	DryRun        bool
}

func main() {
	log.Println("Starting file classifier...")
	args := new(Args)
	c := new(Configuration)

	flag.StringVar(&args.Configuration, "conf", "config.yml", "Configuration file to load")
	flag.BoolVar(&args.DryRun, "dry-run", false, "Don't execute actions, just display what would happen")
	flag.Parse()

	dryRun = args.DryRun

	if dryRun {
		log.Println("Running in dry-run mode")
	}

	err := configor.Load(c, args.Configuration)

	if err != nil || len(c.SearchOperations) == 0 {
		log.Panicf("Impossible to load configuration file. %v", err)
	}

	ch := make(chan UnanalysedFile)
	ch2 := make(chan AnalysedFile)

	executor := DefaultExecutor{}

	go analyse(ch, ch2)
	go processFile(executor, ch2)

	search(c.SearchOperations, ch)
	close(ch)
	<-ch2
}
