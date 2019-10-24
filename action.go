package main

import (
	"errors"
	"log"
	"os"
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

func action(f AnalysedFile, actions []Action) error {
	for _, action := range actions {
		switch action.Operation {
		case copyOperation:
			dest := path.Join(action.Destination, path.Base(f.Path))
			log.Printf("Copying file %v to %v", f.Path, dest)
			err := copy.Copy(f.Path, dest)
			if err != nil {
				return err
			}
		case moveOperation:
			dest := path.Join(action.Destination, path.Base(f.Path))
			log.Printf("Moving file %v to %v", f.Path, dest)
			err := copy.Copy(f.Path, dest)
			if err != nil {
				return err
			}
			err = os.Remove(f.Path)
			if err != nil {
				return err
			}
		default:
			return errors.New("Error : Unknown action")
		}
	}
	return nil
}
