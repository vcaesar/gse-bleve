package main

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve/v2"
	gseb "github.com/vcaesar/gse-bleve"
)

func main() {
	opt := gseb.Option{
		Index: "test.blv",
		// Dicts: "embed, zh",
		Dicts: "embed, ja",
		Stop:  "",
		Opt:   "search-hmm",
		Trim:  "trim",
	}

	index, err := gseb.New(opt)
	if err != nil {
		fmt.Println("new mapping error is: ", err)
		return
	}

	text := `見解では、謙虚なヴォードヴィリアンのベテランは、運命の犠牲者と悪役の両方の変遷として代償を払っています`
	err = index.Index("1", text)

	index.Index("13", "間違っていたのは俺じゃない、世界の方だ")
	index.Index("3", text+"浮き沈み")
	index.Index("10", text+"両方の変遷")
	index.Index("4", `In view, a humble vaudevillian veteran cast vicariously as both victim and villain vicissitudes of fate.`)
	index.Index("2", `It's difficult to understand the sum of a person's life.`)
	if err != nil {
		fmt.Println("index error: ", err)
	}

	query := "運命の犠牲者"
	req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(query))

	req.Highlight = bleve.NewHighlight()
	req.Size = 20
	res, err := index.Search(req)
	fmt.Println("res: ", res, err, res.Hits)

	os.RemoveAll("test.blv")
}
