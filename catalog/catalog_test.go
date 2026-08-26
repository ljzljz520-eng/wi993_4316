package catalog

import (
	"coffeeware/model"
	"coffeeware/store"
	"path/filepath"
	"testing"
)

func TestCatalogSearch(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := New(s)
	_ = c.Register(model.NewRecord("a1", "Gooseneck", "kettle", 90, 2))
	rs, e := c.Search("neck", "", "")
	if e != nil || len(rs) != 1 {
		t.Fatal(e)
	}
}
