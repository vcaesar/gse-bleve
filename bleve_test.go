package gsebleve

import (
	"os"
	"testing"

	"github.com/blevesearch/bleve/v2"
	"github.com/vcaesar/tt"
)

func TestGse(t *testing.T) {
	indexName := "test.blv"
	os.RemoveAll(indexName)

	opt := Option{
		Index: indexName,
		Dicts: "embed, zh",
		Opt:   "serch-hmm",
	}
	index, err := New(opt)
	tt.Nil(t, err)
	err = index.Index("10", "多伦多 the CN Tower, 多伦多悬崖公园")
	tt.Nil(t, err)
	err = index.Index("1", "西雅图太空针")
	tt.Nil(t, err)

	qry := "多伦多"
	req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(qry))
	req.Highlight = bleve.NewHighlight()
	res, err := index.Search(req)
	tt.Nil(t, err)
	tt.Equal(t, 1, res.Total)

	tt.Equal(t, 1, res.Hits.Len())
	tt.Equal(t, "map[]", res.Hits[0].Fragments)
	tt.Equal(t, "[]", res.Request.Fields)
	tt.Equal(t, indexName, res.Hits[0].Index)

	os.RemoveAll(indexName)
}
