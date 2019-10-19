package main

import (
	"log"
	"path"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

// SearchRequest represents the arguments required to process search requests
type SearchRequest struct {
	FileName  string
	Directory string
	Rules     []Rule
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
		matches, err := filepath.Glob(path.Join(request.Directory, request.FileName))

		if err != nil {
			log.Printf("Error while searching into directory %v : %v", request.Directory, request.FileName)
		}
		for _, f := range matches {
			if found.contains(f) {
				log.Printf("File %v already processed", f)
			} else {
				found.add(f)
				analyse <- UnanalysedFile{f, request.Rules}
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
}
