// Copyright 2016 Evans. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gsebleve

import (
	"strings"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/go-ego/gse"
)

const (
	TokenName = "gse"
)

// GseCut gse cut token structure
type GseCut struct {
	seg *gse.Segmenter
	// stop string
	opt  string
	trim string
}

// NewGseCut create a gse cut tokenizer
func NewGse(dicts, stop, opt, trim string) (*GseCut, error) {
	var (
		seg gse.Segmenter
		err error
	)

	seg.SkipLog = true
	if dicts == "" {
		dicts = "zh"
	}

	if strings.Contains(dicts, "emend") {
		dicts = strings.Replace(dicts, "emend, ", "", 1)
		err = seg.LoadDictEmbed(dicts)
	} else {
		err = seg.LoadDict(dicts)
	}

	if stop != "" {
		if strings.Contains(stop, "emend") {
			stop = strings.Replace(stop, "emend, ", "", 1)
			seg.LoadStopEmbed(stop)
		} else {
			seg.LoadStop(stop)
		}
	}
	return &GseCut{&seg, opt, trim}, err
}

// Trim trim the unused token string
func (c *GseCut) Trim(s []string) []string {
	if c.trim == "symbol" {
		return c.seg.TrimSymbol(s)
	}

	if c.trim == "punct" {
		return c.seg.TrimPunct(s)
	}

	if c.trim == "trim" {
		return c.seg.Trim(s)
	}

	return s
}

// Cut option the gse cut mode
func (c *GseCut) Cut(text string, opt string) []string {
	if c.trim == "html" {
		return c.seg.CutTrimHtml(text)
	}

	if c.trim == "url" {
		return c.seg.CutUrl(text)
	}

	if opt == "search-hmm" {
		return c.seg.CutSearch(text, true)
	}
	if opt == "search" {
		return c.seg.CutSearch(text)
	}

	if opt == "search-dag" {
		return c.seg.CutSearch(text, false)
	}

	if opt == "all" {
		return c.seg.CutAll(text)
	}

	if opt == "hmm" {
		return c.seg.Cut(text, true)
	}

	if opt == "dag" {
		return c.seg.Cut(text, false)
	}

	return c.seg.Cut(text)
}

// Tokenize cut the text to bleve token stream
func (c *GseCut) Tokenize(text []byte) analysis.TokenStream {
	result := make(analysis.TokenStream, 0)
	cuts := c.Trim(c.Cut(string(text), c.opt))
	// fmt.Println("cuts: ", cuts)
	azs := c.seg.Analyze(cuts)
	for _, az := range azs {
		token := analysis.Token{
			Term:     []byte(az.Text),
			Start:    az.Start,
			End:      az.End,
			Position: az.Position,
			Type:     analysis.Ideographic,
		}
		result = append(result, &token)
	}
	return result
}

func tokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	dicts, ok := config["dicts"].(string)
	if !ok {
		dicts = ""
	}
	stop, ok := config["stop"].(string)
	if !ok {
		stop = ""
	}

	opt, ok := config["opt"].(string)
	if !ok {
		opt = ""
	}

	trim, ok := config["trim"].(string)
	if !ok {
		trim = ""
	}

	return NewGse(dicts, stop, opt, trim)
}

func init() {
	registry.RegisterTokenizer(TokenName, tokenizerConstructor)
}
