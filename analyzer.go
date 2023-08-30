// Copyright 2016 Evans. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gsebleve

import (
	"errors"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

func NewAnalyzer(config map[string]interface{}, cache *registry.Cache) (analysis.Analyzer, error) {
	tokenizerName, ok := config["tokenizer"].(string)
	if !ok {
		return nil, errors.New("must have tokenizer")
	}

	tokenizer, err := cache.TokenizerNamed(tokenizerName)
	if err != nil {
		return nil, err
	}

	az := &analysis.DefaultAnalyzer{Tokenizer: tokenizer}
	return az, nil
}

func init() {
	registry.RegisterAnalyzer(TokenName, NewAnalyzer)
	registry.RegisterAnalyzer(SeparateName, NewAnalyzer)
}
