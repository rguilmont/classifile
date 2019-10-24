package main

import (
	"fmt"
	"log"
	"regexp"
)

// Rule is a list of Match that neeeds to be validated
type Rule struct {
	Conditions []Match
	Actions    []Action
}

// Match is verified against a file content, title, metadata...
type Match struct {
	Elem     string
	Matches  string
	Expected bool
}

func validateRule(r Rule) bool {
	for _, m := range r.Conditions {
		if m.Elem != "content" &&
			m.Elem != "metadata" &&
			m.Elem != "type" &&
			m.Elem != "path" {
			return false
		}
	}
	return true
}

func elements(elem string, f AnalysedFile) []string {
	e := []string{}
	if elem == "content" {
		e = append(e, f.TextContent)
	} else if elem == "metadata" {
		for _, meta := range f.Meta {
			fmt.Println(meta)
			e = append(e, meta)
		}
	} else if elem == "type" {
		e = append(e, f.Type)
	} else if elem == "path" {
		e = append(e, f.Path)
	}
	return e
}

func processRule(r Rule, f AnalysedFile) bool {

	for _, m := range r.Conditions {
		if !validateRule(r) {
			log.Printf("Invalid rule %v", m)
			return false
		}

		e := elements(m.Elem, f)
		for _, content := range e {
			res, err := regexp.MatchString(m.Matches, content)
			if err != nil {
				log.Printf("Error while processing match %v : %v", m, err)
				return false
			}
			if res != m.Expected {
				log.Printf("PROCESSING RULE %v - %v NOT TRIGGERED", r, f.Path)
				return false
			}
		}
	}
	log.Printf("PROCESSING RULE %v - %v TRIGGERED", r, f.Path)
	return true
}

func processFile(process chan AnalysedFile) {
	for f := range process {
		hasMatched := false
		for _, rule := range f.Rules {
			if processRule(rule, f) {
				err := action(f, rule.Actions)
				if err != nil {
					log.Printf("Error while proceding to actions %v : %v", f.Actions, err)
				}
				hasMatched = true
			}
			if hasMatched {
				break
			}
		}
	}
}
