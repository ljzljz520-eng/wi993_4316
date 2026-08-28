package store

import (
	"coffeeware/model"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := model.NewRecord("persist-1", "Dripper", "brewing", 40, 1)
	if e = s.PutRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.GetRecord("persist-1")
	if e != nil || got.Name != r.Name {
		t.Fatalf("reopen: %v", e)
	}
}
