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

func run(args Args) {
	conf := new(Configuration)
	if args.DryRun {
		log.Println("Running in dry-run mode.")
	}

	err := configor.Load(conf, args.Configuration)

	if err != nil || len(conf.SearchOperations) == 0 {
		log.Panicf("Impossible to load configuration file. %v", err)
	}

	ch := make(chan UnanalysedFile)
	ch2 := make(chan AnalysedFile)

	executor := DefaultExecutor{}

	go analyse(ch, ch2)
	go processFile(executor, ch2)

	search(conf.SearchOperations, ch)
	close(ch)
	<-ch2
}

func main() {
	log.Println("Starting file classifier...")
	args := new(Args)
	flag.StringVar(&args.Configuration, "conf", "config.yml", "Configuration file to load")
	flag.BoolVar(&args.DryRun, "dry-run", false, "Don't execute actions, just display what would happen")
	flag.Parse()

	run(*args)
}
