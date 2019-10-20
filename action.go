package main

import (
	"log"
	"path"

	"github.com/otiai10/copy"
)

// Action represent an action to a file when matching a rule. They can be chained.
type Action struct {
	Operation   string
	Destination string
}

const (
	moveOperation string = "move"
	copyOperation string = "copy"
)

func action(f AnalysedFile) {
	for _, action := range f.Actions {
		switch action.Operation {
		case copyOperation:
			dest := path.Join(action.Destination, path.Base(f.Path))
			log.Printf("Copying file %v to %v", f.Path, dest)
			err := copy.Copy(f.Path, dest)
			if err != nil {
				log.Printf("Error while copying file %v : %v", f.Path, err)
			}
		}
	}
}
