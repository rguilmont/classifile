package main

import (
	"fmt"
	"regexp"

	log "github.com/sirupsen/logrus"
)

// Rule is a list of Match that neeeds to be validated
type Rule struct {
	Conditions []Match  `required:"true"`
	Actions    []Action `required:"true"`
}

// Match is verified against a file content, title, metadata...
type Match struct {
	Elem     string `required:"true"`
	Matches  string `required:"true"`
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
			log.Warnf("Invalid rule %v", m)
			return false
		}

		e := elements(m.Elem, f)
		for _, content := range e {
			res, err := regexp.MatchString(m.Matches, content)
			if err != nil {
				log.Warnf("Error while processing match %v : %v", m, err)
				return false
			}
			if res != m.Expected {
				log.Debugf("PROCESSING RULE %v - %v NOT TRIGGERED", r, f.Path)
				return false
			}
		}
	}
	log.Debugf("PROCESSING RULE %v - %v TRIGGERED", r, f.Path)
	return true
}

func processFile(exec Executor, process chan AnalysedFile) {
	for f := range process {
		hasMatched := false
		for _, rule := range f.Rules {
			if processRule(rule, f) {
				err := action(exec, f, rule.Actions)
				if err != nil {
					log.Warnf("Error while proceding to actions %v : %v", rule.Actions, err)
				}
				hasMatched = true
			}
			if hasMatched {
				break
			}
		}
	}
}
