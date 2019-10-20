package main

import (
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

	analysed.Meta = res.Meta
	analysed.Path = f.Path
	analysed.TextContent = res.Body
	analysed.Type = fileType
	analysed.Rules = f.Rules
	analysed.Actions = f.Actions

	return analysed, nil
}
