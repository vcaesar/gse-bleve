package gsebleve

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/blevesearch/bleve/v2"
	"github.com/vcaesar/tt"
)

type data struct {
	Id       string
	Title    string
	Content  string
	UpdateAt *time.Time
}

var text = "Seattle space needle, 多伦多 the CN Tower, 多伦多悬崖公园"

func TestGse(t *testing.T) {
	indexName := "test.blv"
	os.RemoveAll(indexName)

	opt := Option{
		Index: indexName,
		// Dicts: "embed, ja",
		Dicts: "embed, zh",
		Opt:   "search-hmm",
		Trim:  "trim",
		Stop:  "",
	}

	// index, err := New(opt)
	//
	m1, err := NewMapping(opt)
	tt.Nil(t, err)
	docMap := bleve.NewDocumentMapping()
	m2 := NewTextMap()
	m2.IncludeTermVectors = true
	docMap.AddFieldMappingsAt("title", m2)
	m1.AddDocumentMapping("text1", docMap)

	index, err := bleve.New(opt.Index, m1)
	tt.Nil(t, err)

	err = index.Index("10", text)
	tt.Nil(t, err)
	err = index.Index("1", "西雅图太空针")
	tt.Nil(t, err)

	for i := 0; i < 10; i++ {
		t1 := time.Now()
		d1 := data{
			Id:       "10" + fmt.Sprint(i),
			Title:    text + fmt.Sprint(i),
			Content:  text,
			UpdateAt: &t1,
		}

		err = index.Index(d1.Id, d1)
		tt.Nil(t, err)
	}

	qry := "Seattle"
	req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(qry))
	req.Size = 20
	req.Highlight = bleve.NewHighlight()
	res, err := index.Search(req)
	tt.Nil(t, err)
	tt.Equal(t, 11, res.Total)
	fmt.Println("res: ", res)

	tt.Equal(t, 11, res.Hits.Len())
	tt.Equal(t, "map[]", res.Hits[0].Fragments)
	tt.Equal(t, "[]", res.Request.Fields)
	tt.Equal(t, indexName, res.Hits[0].Index)

	os.RemoveAll(indexName)
}
