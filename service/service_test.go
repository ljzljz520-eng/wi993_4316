package service

import (
	"coffeeware/model"
	"coffeeware/store"
	"context"
	"path/filepath"
	"testing"
)

func TestServiceChain(t *testing.T) {
	s0, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s0.Close()
	s := New(s0)
	if e := s.Register(model.NewRecord("svc1", "Scale", "tools", 10, 2)); e != nil {
		t.Fatal(e)
	}
	if e := s.Submit(context.Background(), "svc1"); e != nil {
		t.Fatal(e)
	}
	if _, e := s.Metrics(); e != nil {
		t.Fatal(e)
	}
}
