package main

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/otiai10/copy"
)

// Executor describes the actions implementation that can be applied to a file
type Executor interface {
	copy(string, string) error
	move(string, string) error
}

// DefaultExecutor applies real actions to files, unlike other implementations used for tests
//  or for dry run.
type DefaultExecutor struct{}

func (d DefaultExecutor) copy(src string, dest string) error {
	return copy.Copy(src, dest)
}

func (d DefaultExecutor) move(src string, dest string) error {
	err := copy.Copy(src, dest)
	if err == nil {
		return os.Remove(src)
	}
	return err
}

// Action represent an action to a file when matching a rule. They can be chained.
type Action struct {
	Operation   string
	Destination string
}

const (
	moveOperation string = "move"
	copyOperation string = "copy"
)

func action(exec Executor, f AnalysedFile, actions []Action) error {
	for _, action := range actions {
		switch action.Operation {
		case copyOperation:
			dest := path.Join(action.Destination, path.Base(f.Path))
			log.Printf("Copying file %v to %v", f.Path, dest)

			if !dryRun {
				err := exec.copy(f.Path, dest)
				if err != nil {
					return err
				}
			}
		case moveOperation:
			dest := path.Join(action.Destination, path.Base(f.Path))
			log.Printf("Moving file %v to %v", f.Path, dest)
			if !dryRun {
				err := exec.move(f.Path, dest)
				if err != nil {
					return err
				}
			}
		default:
			return errors.New("Error : Unknown action")
		}
	}
	return nil
}
