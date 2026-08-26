package workflow

import (
	"coffeeware/catalog"
	"coffeeware/model"
	"context"
	"fmt"
)

type Pipeline struct {
	cat *catalog.Catalog
	mon *Monitor
}

func NewPipeline(c *catalog.Catalog, m *Monitor) *Pipeline { return &Pipeline{cat: c, mon: m} }
func (p *Pipeline) Receive(id, name, cat string, price float64, stock int) error {
	return p.cat.Register(model.NewRecord(id, name, cat, price, stock))
}
func (p *Pipeline) Process(ctx context.Context, id string) error {
	if _, e := p.cat.Publish(id); e != nil {
		return e
	}
	return p.mon.Start(ctx, id, 3)
}
func (p *Pipeline) Archive(id string) error { _, e := p.cat.Archive(id); return e }
func (p *Pipeline) Track(id string) (string, error) {
	s := p.mon.State(id)
	if s == "" {
		return "", fmt.Errorf("unknown task")
	}
	return s, nil
}
func (p *Pipeline) Reconcile(id string) error {
	r, e := p.cat.Find(id)
	if e != nil {
		return e
	}
	if r.Status == "published" && p.mon.State(id) == "completed" {
		return nil
	}
	return fmt.Errorf("not reconciled")
}
