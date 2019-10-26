package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/configor"
)

// Global flags
var dryRun = false

// Export variables for version
var (
	GitSummary string
	BuildDate  string
)

// Configuration is the configuration format
type Configuration struct {
	SearchOperations []SearchRequest `required:"true"`
}

// Args are startup arguments
type Args struct {
	Configuration string
	DryRun        bool
	Version       bool
}

func run(args Args) {
	if args.Version {
		fmt.Printf("Version %v  - build at %v\n", GitSummary, BuildDate)
		os.Exit(0)
	}

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
	args := new(Args)
	flag.StringVar(&args.Configuration, "conf", "config.yml", "Configuration file to load")
	flag.BoolVar(&args.DryRun, "dry-run", false, "Don't execute actions, just display what would happen")
	flag.BoolVar(&args.Version, "version", false, "display version")
	flag.Parse()

	run(*args)
}
