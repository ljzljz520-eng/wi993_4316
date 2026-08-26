package catalog

import (
	"coffeeware/model"
	"coffeeware/store"
	"path/filepath"
	"testing"
)

func TestBulkImport(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r := ImportResult{}
	r = New(s).Import([]model.Record{model.NewRecord("bulk1", "Scale", "tools", 10, 1)})
	if r.Accepted != 1 {
		t.Fatal(r)
	}
}
