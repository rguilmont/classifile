package main

import (
	"errors"

	"code.sajari.com/docconv"
)

type UnanalysedFile struct {
	Path    string
	Rules   []Rule
	Actions []Action
}

// AnalysedFile represent a file once analysed, that can then be processed
type AnalysedFile struct {
	TextContent string
	Meta        map[string]string
	Type        string
	Path        string
	Rules       []Rule
	Actions     []Action
}

// Convert a file, depending on its type
func convert(f UnanalysedFile, fileType string) (AnalysedFile, error) {

	analysed := AnalysedFile{}

	res, err := docconv.ConvertPath(f.Path)
	if err != nil {
		return analysed, err
	}
	// When docconv can't convert a file, instead of returning an error, it returns
	//  an empty Body and empty Metadata
	if res.Body == "" && len(res.Meta) == 0 {
		return analysed, errors.New("Can't convert file")
	}

	analysed.Meta = res.Meta
	analysed.Path = f.Path
	analysed.TextContent = res.Body
	analysed.Type = fileType
	analysed.Rules = f.Rules
	analysed.Actions = f.Actions

	return analysed, nil
}
