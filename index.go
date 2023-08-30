// Copyright 2016 Evans. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gsebleve

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

// Option gse bleve option structure
type Option struct {
	Index                  string
	Dicts, Stop, Opt, Trim string
	Alpha                  bool
	Name, Sep              string
}

// NewMappingSep new separator mapping
func NewMappingSep(sep string, trim ...string) (*mapping.IndexMappingImpl, error) {
	mapping1 := bleve.NewIndexMapping()
	trimOpt := ""
	if len(trim) > 0 {
		trimOpt = trim[0]
	}

	err := mapping1.AddCustomTokenizer(SeparateName, map[string]interface{}{
		"type": SeparateName,
		"sep":  sep,
		"trim": trimOpt,
	})
	if err != nil {
		return nil, err
	}

	err = mapping1.AddCustomAnalyzer(SeparateName, map[string]interface{}{
		"type":      SeparateName,
		"tokenizer": SeparateName,
	})
	if err != nil {
		return nil, err
	}

	mapping1.DefaultAnalyzer = SeparateName
	return mapping1, nil
}

// NewMapping new bleve index mapping
func NewMapping(opt Option) (*mapping.IndexMappingImpl, error) {
	mapping1 := bleve.NewIndexMapping()

	err := mapping1.AddCustomTokenizer(TokenName, map[string]interface{}{
		"type":  TokenName,
		"dicts": opt.Dicts,
		"stop":  opt.Stop,
		"opt":   opt.Opt,
		"trim":  opt.Trim,
		"alpha": opt.Alpha,
	})

	if err != nil {
		return nil, err
	}

	err = mapping1.AddCustomAnalyzer(TokenName, map[string]interface{}{
		"type":      TokenName,
		"tokenizer": TokenName,
	})

	if err != nil {
		return nil, err
	}

	mapping1.DefaultAnalyzer = TokenName
	return mapping1, nil
}

// New new bleve index
func New(opt Option) (bleve.Index, error) {
	var (
		mapping1 *mapping.IndexMappingImpl
		err      error
	)
	if opt.Name == "sep" {
		mapping1, err = NewMappingSep(opt.Sep, opt.Trim)
	} else {
		mapping1, err = NewMapping(opt)
	}
	if err != nil {
		return nil, err
	}

	return bleve.New(opt.Index, mapping1)
}

// NewMem new bleve index only memory
func NewMem(opt Option) (bleve.Index, error) {
	mapping1, err := NewMapping(opt)
	if err != nil {
		return nil, err
	}

	return bleve.NewMemOnly(mapping1)
}

// NewDoc new bleve index document mapping
func NewDoc() *mapping.DocumentMapping {
	return bleve.NewDocumentMapping()
}

// NewTextMap new text field mapping with gse
func NewTextMap() *mapping.FieldMapping {
	return &mapping.FieldMapping{
		Type:         "text",
		Analyzer:     "gse",
		Store:        true,
		Index:        true,
		IncludeInAll: true,
		DocValues:    true,
	}
}

// NewSepMap new text field mapping with sep
func NewSepMap() *mapping.FieldMapping {
	return &mapping.FieldMapping{
		Type:         "text",
		Analyzer:     "sep",
		Store:        true,
		Index:        true,
		IncludeInAll: true,
		DocValues:    true,
	}
}
