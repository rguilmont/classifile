package main

import (
	"testing"
	"time"
)

func TestAbsentPathSearchRequest(t *testing.T) {
	s := SearchRequest{
		".+",
		"this is not an existing path, should not match",
		[]Rule{},
	}

	matched := 0

	c := make(chan UnanalysedFile)
	end := make(chan bool)

	go func() {
		for range c {
			matched++
		}
		close(end)
	}()

	search([]SearchRequest{s}, c)
	close(c)
	<-end
	if matched != 0 {
		t.Errorf("Expecting no result from search, got at least %v results", matched)
	}
}

func TestValidSearchRequest(t *testing.T) {
	s := SearchRequest{
		".+.pdf",
		"./test_assets",
		[]Rule{},
	}

	var matched = 0

	c := make(chan UnanalysedFile)
	end := make(chan bool)

	go func() {
		for range c {
			matched++
		}
		close(end)
	}()

	search([]SearchRequest{s}, c)
	close(c)
	<-end

	if matched != 1 {
		t.Errorf("Expecting 1 result for search %v, got %v", s, matched)
	}
}

func TestMultipleSearchRequest(t *testing.T) {
	s := SearchRequest{
		".+.pdf",
		"./test_assets",
		[]Rule{},
	}

	var matched = 0

	c := make(chan UnanalysedFile)
	end := make(chan bool)

	go func() {
		for range c {
			matched++
		}
		close(end)
	}()

	search([]SearchRequest{s, s, s, s, s}, c)
	close(c)
	<-end

	if matched != 1 {
		t.Errorf("Expecting 1 result for search %v, got %v", s, matched)
	}
}

func TestAnalyseRequest(t *testing.T) {

	c1 := make(chan UnanalysedFile)
	c2 := make(chan AnalysedFile)

	defer close(c2)
	go analyse(c1, c2)
	go func() {
		c1 <- UnanalysedFile{
			"./test_assets/file-sample_100kB.docx",
			[]Rule{},
		}
	}()

	select {
	case analysed := <-c2:
		if analysed.Type != "application/vnd.openxmlformats-officedocument.wordprocessingml.document" ||
			analysed.TextContent == "" {
			t.Errorf("Error while getting analysed file %v", analysed)
		}
	case <-time.After(10 * time.Second):
		t.Errorf("Expecting analysed file, got timeout")
	}
}
