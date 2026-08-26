package workflow

import (
	"coffeeware/model"
	"coffeeware/store"
	"context"
	"sync"
	"time"
)

type Monitor struct {
	st    *store.Store
	mu    sync.Mutex
	state map[string]string
}

func NewMonitor(st *store.Store) *Monitor { return &Monitor{st: st, state: map[string]string{}} }
func (m *Monitor) Start(ctx context.Context, id string, workers int) error {
	m.mu.Lock()
	m.state[id] = "running"
	m.mu.Unlock()
	if workers < 1 {
		workers = 1
	}
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(n int) { defer wg.Done(); m.child(ctx, id, n) }(i)
	}
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-ctx.Done():
		m.mu.Lock()
		m.state[id] = "cancelled"
		m.mu.Unlock()
		return model.ErrCancelled
	case <-done:
		m.mu.Lock()
		m.state[id] = "completed"
		m.mu.Unlock()
		return nil
	}
}
func (m *Monitor) child(ctx context.Context, id string, n int) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Millisecond * 2)
		_ = ctx
	}
	m.st.PutEvent(model.NewEvent(id+"-"+string(rune(n)), id, "child", "system", "finished"))
}
func (m *Monitor) State(id string) string { m.mu.Lock(); defer m.mu.Unlock(); return m.state[id] }
func (m *Monitor) Cancel(id string) {
	m.mu.Lock()
	if m.state[id] == "running" {
		m.state[id] = "cancelling"
	}
	m.mu.Unlock()
}
