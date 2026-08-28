package workflow

import (
	"coffeeware/catalog"
	"coffeeware/store"
	"context"
	"path/filepath"
	"testing"
	"time"
)

func TestWorkflowThree(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	p := NewPipeline(catalog.New(s), NewMonitor(s))
	_ = p.Receive("w3", "Basket", "storage", 15, 2)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if e := p.Process(ctx, "w3"); e == nil {
		t.Fatal("expected cancellation")
	}
}
func TestBusinessChain29(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	p := NewPipeline(catalog.New(s), NewMonitor(s))
	_ = p.Receive("cancel-29", "Server", "kettle", 20, 1)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	e := p.Process(ctx, "cancel-29")
	if e == nil {
		t.Fatal("expected cancellation")
	}
	time.Sleep(time.Millisecond * 10)
	if p.mon.State("cancel-29") == "running" {
		t.Fatal("task remains running")
	}
	time.Sleep(time.Millisecond * 60)
	if n, _ := s.Count("events"); n != 0 {
		t.Fatalf("cancelled task emitted %d events", n)
	}
}
