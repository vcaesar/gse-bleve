# gse-bleve

[![Build Status](https://github.com/vcaesar/gse-bleve/workflows/Go/badge.svg)](https://github.com/vcaesar/gse-bleve/commits/master)
[![Build Status](https://travis-ci.org/vcaesar/gse-bleve.svg)](https://travis-ci.org/vcaesar/gse-bleve)
[![CircleCI Status](https://circleci.com/gh/vcaesar/gse-bleve.svg?style=shield)](https://circleci.com/gh/vcaesar/gse-bleve)
[![codecov](https://codecov.io/gh/vcaesar/gse-bleve/branch/master/graph/badge.svg)](https://codecov.io/gh/vcaesar/gse-bleve)
[![Go Report Card](https://goreportcard.com/badge/github.com/vcaesar/gse-bleve)](https://goreportcard.com/report/github.com/vcaesar/gse-bleve)
[![GoDoc](https://godoc.org/github.com/vcaesar/gse-bleve?status.svg)](https://godoc.org/github.com/vcaesar/gse-bleve)
[![Release](https://github-release-version.herokuapp.com/github/vcaesar/gse-bleve/release.svg?style=flat)](https://github.com/vcaesar/gse-bleve/releases/latest)


## Use 

```go
package main

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve/v2"
	gse "github.com/vcaesar/gse-bleve"
)

func main() {
	opt := gse.Option{
		Index: "test.blv",
		Dicts: "embed, zh", Stop: "",
		Opt: "search-hmm", Trim: "trim"}

	index, err := gse.New(opt)
	if err != nil {
		fmt.Println("new mapping error is: ", err)
	}

	text := `他在命运的沉浮中随波逐流, 扮演着受害与加害者的双重角色`
	err = index.Index("1", text)
	index.Index("3", text+"沉浮")
	index.Index("4", `In view, a humble vaudevillian veteran cast vicariously as both victim and villain vicissitudes of fate.`)
	index.Index("2", `It's difficult to understand the sum of a person's life.`)
	if err != nil {
		fmt.Println("index error: ", err)
	}

	query := "命运的沉浮"
	req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(query))
	req.Highlight = bleve.NewHighlight()
	res, err := index.Search(req)
	fmt.Println(res, err)

	os.RemoveAll("test.blv")
}
```