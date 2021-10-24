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
		Dicts: "embed, zh",
		// Stop: "",
		Opt: "dag",
		// Trim: "trim",
	}

	index, err := gse.New(opt)
	if err != nil {
		fmt.Println("new mapping error is: ", err)
		return
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
