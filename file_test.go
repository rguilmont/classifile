package main

import (
	"strings"
	"testing"
)

// All tests files are downloaded from https://file-examples.com/index.php/sample-documents-download/

const validationString = "Vestibulum neque massa, scelerisque sit amet ligula eu, congue molestie mi"

func TestUnknownFileConvertion(t *testing.T) {

	f := UnanalysedFile{
		"./test_assets/invalid.exe",
		[]Rule{},
		[]Action{},
	}

	af, err := convert(f, "")
	//t.Log(err)
	if err == nil {
		t.Errorf("Was expecting an error, got a valid file : %v - %v ", af.TextContent, af.Meta)
	}
}

func TestNotExistingFileConvertion(t *testing.T) {
	f := UnanalysedFile{
		"./test_assets/not_existing.doc",
		[]Rule{},
		[]Action{},
	}

	af, err := convert(f, "")
	//t.Log(err)
	if err == nil {
		t.Errorf("Was expecting an error, got a valid file : %v - %v ", af.TextContent, af.Meta)
	}
}

func TestDocFileConvertion(t *testing.T) {
	f := UnanalysedFile{
		"./test_assets/file-sample_100kB.doc",
		[]Rule{},
		[]Action{},
	}

	af, err := convert(f, "")
	//t.Log(err)
	if err != nil {
		t.Errorf("Was expecting an error, got a valid file : %v - %v ", af.TextContent, af.Meta)
	}
	if !strings.Contains(af.TextContent, validationString) {
		t.Errorf("Can't find Valid string in file")
	}
}

func TestDocxFileConvertion(t *testing.T) {
	f := UnanalysedFile{
		"./test_assets/file-sample_100kB.docx",
		[]Rule{},
		[]Action{},
	}

	af, err := convert(f, "")
	//t.Log(err)
	if err != nil {
		t.Errorf("Was expecting an error, got a valid file : %v - %v ", af.TextContent, af.Meta)
	}
	if !strings.Contains(af.TextContent, validationString) {
		t.Errorf("Can't find Valid string in file")
	}
}

func TestPdfFileConvertion(t *testing.T) {
	f := UnanalysedFile{
		"./test_assets/file-sample_150kB.pdf",
		[]Rule{},
		[]Action{},
	}

	af, err := convert(f, "")
	//t.Log(err)
	if err != nil {
		t.Errorf("Was expecting a valid file, got an error: %v", err)
	}
	if !strings.Contains(af.TextContent, validationString) {
		t.Errorf("Can't find Valid string in file")
	}
}

func TestOdtFileConvertion(t *testing.T) {
	f := UnanalysedFile{
		"./test_assets/file-sample_100kB.odt",
		[]Rule{},
		[]Action{},
	}

	af, err := convert(f, "")
	//t.Log(err)
	if err != nil {
		t.Errorf("Was expecting a valid file, got an error: %v", err)
	}
	if !strings.Contains(af.TextContent, validationString) {
		t.Errorf("Can't find Valid string in file")
	}
}
