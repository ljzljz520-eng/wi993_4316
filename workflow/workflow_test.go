package workflow

import (
	"coffeeware/catalog"
	"coffeeware/store"
	"context"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	p := NewPipeline(catalog.New(s), NewMonitor(s))
	if e := p.Receive("w1", "Moka", "brewing", 25, 2); e != nil {
		t.Fatal(e)
	}
	if e := p.Process(context.Background(), "w1"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	p := NewPipeline(catalog.New(s), NewMonitor(s))
	_ = p.Receive("w2", "Tamper", "tools", 12, 2)
	_ = p.Process(context.Background(), "w2")
	if e := p.Archive("w2"); e != nil {
		t.Fatal(e)
	}
}
