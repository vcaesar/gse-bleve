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
	})
	if err != nil {
		return mapping, err
	}

	err = mapping.AddCustomAnalyzer(TokenName, map[string]interface{}{
		"type":      TokenName,
		"tokenizer": TokenName,
	})

	if err != nil {
		return mapping, err
	}

	mapping.DefaultAnalyzer = TokenName
	return mapping, nil
}

// New new bleve index
func New(opt Option) (bleve.Index, error) {
	mapping, err := NewMapping(opt)
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
