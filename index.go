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
	mapping := bleve.NewIndexMapping()
	trimOpt := ""
	if len(trim) > 0 {
		trimOpt = trim[0]
	}

	err := mapping.AddCustomTokenizer(SeparateName, map[string]interface{}{
		"type": SeparateName,
		"sep":  sep,
		"trim": trimOpt,
	})
	if err != nil {
		return nil, err
	}

	err = mapping.AddCustomAnalyzer(SeparateName, map[string]interface{}{
		"type":      SeparateName,
		"tokenizer": SeparateName,
	})
	if err != nil {
		return nil, err
	}

	mapping.DefaultAnalyzer = SeparateName
	return mapping, nil
}

// NewMapping new bleve index mapping
func NewMapping(opt Option) (*mapping.IndexMappingImpl, error) {
	mapping := bleve.NewIndexMapping()

	err := mapping.AddCustomTokenizer(TokenName, map[string]interface{}{
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

	err = mapping.AddCustomAnalyzer(TokenName, map[string]interface{}{
		"type":      TokenName,
		"tokenizer": TokenName,
	})

	if err != nil {
		return nil, err
	}

	mapping.DefaultAnalyzer = TokenName
	return mapping, nil
}

// New new bleve index
func New(opt Option) (bleve.Index, error) {
	var (
		mapping *mapping.IndexMappingImpl
		err     error
	)
	if opt.Name == "sep" {
		mapping, err = NewMappingSep(opt.Sep, opt.Trim)
	} else {
		mapping, err = NewMapping(opt)
	}
	if err != nil {
		return nil, err
	}

	return bleve.New(opt.Index, mapping)
}

// NewMem new bleve index only memory
func NewMem(opt Option) (bleve.Index, error) {
	mapping, err := NewMapping(opt)
	if err != nil {
		return nil, err
	}

	return bleve.NewMemOnly(mapping)
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
