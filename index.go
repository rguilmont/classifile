package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gabriel-vasile/mimetype"
)

// SearchRequest represents the arguments required to process search requests
type SearchRequest struct {
	FileName  string
	Directory string
	Rules     []Rule
	Actions   []Action
}

// Internal type to have a set like structure
type set struct {
	m map[string]bool
}

func newset() *set {
	s := &set{}
	s.m = make(map[string]bool)
	return s
}

// Remove a set
func (s *set) add(value string) {
	s.m[value] = true
}

func (s *set) remove(value string) {
	delete(s.m, value)
}

func (s *set) contains(value string) bool {
	_, c := s.m[value]
	return c
}

// search documents, and send them in analyse channel to be processed
func search(requests []SearchRequest, analyse chan UnanalysedFile) {

	found := newset()

	for _, request := range requests {

		matchingFiles := newset()
		err := filepath.Walk(request.Directory, func(p string, f os.FileInfo, err error) error {
			if err != nil {
				log.Printf("Error while searching on directory %v : %v", p, err)
			} else {
				match, err := regexp.MatchString(request.FileName, f.Name())
				if err != nil {
					log.Printf("Error while searchgin with regex %v : %v", request.FileName, err)
				} else {
					if !f.IsDir() && match {
						matchingFiles.add(p)
					}
				}
			}
			return nil
		})

		if err != nil {
			log.Printf("Error while searching into directory %v : %v", request.Directory, request.FileName)
		}
		for f := range matchingFiles.m {
			if found.contains(f) {
				log.Printf("File %v already processed", f)
			} else {
				found.add(f)
				analyse <- UnanalysedFile{f, request.Rules, request.Actions}
			}
		}
	}
}

// analyse a file, usually by reading or converting it, and them send it to
func analyse(analyse chan UnanalysedFile, process chan AnalysedFile) {
	for f := range analyse {
		log.Printf("Processing file %v", f)

		// First get the type of the file
		fileType, _, err := mimetype.DetectFile(f.Path)
		if err != nil {
			log.Printf("Error while detecting type of file %v : %v", f, err)
		} else {
			analysed, err := convert(f, fileType)
			if err != nil {
				log.Printf("Error while converting file %v : %v", f, err)
			} else {
				process <- analysed
			}
		}
	}
	close(process)
}
