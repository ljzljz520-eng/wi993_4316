package store

import (
	"coffeeware/model"
	"path/filepath"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r := model.NewRecord("r1", "Press", "brewing", 30, 3)
	if e = s.PutRecord(r); e != nil {
		t.Fatal(e)
	}
	if _, e = s.GetRecord("r1"); e != nil {
		t.Fatal(e)
	}
}
