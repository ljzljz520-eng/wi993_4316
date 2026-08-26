package service

import (
	"coffeeware/catalog"
	"coffeeware/model"
	"coffeeware/store"
	"coffeeware/workflow"
	"context"
	"fmt"
)

type Service struct {
	Store    *store.Store
	Catalog  *catalog.Catalog
	Pipeline *workflow.Pipeline
}

func New(st *store.Store) *Service {
	c := catalog.New(st)
	m := workflow.NewMonitor(st)
	return &Service{Store: st, Catalog: c, Pipeline: workflow.NewPipeline(c, m)}
}
func (s *Service) Register(r model.Record) error               { return s.Catalog.Register(r) }
func (s *Service) Submit(ctx context.Context, id string) error { return s.Pipeline.Process(ctx, id) }
func (s *Service) Archive(id string) error                     { return s.Pipeline.Archive(id) }
func (s *Service) Query(term string) ([]model.Record, error)   { return s.Catalog.Search(term, "", "") }
func (s *Service) Status(id string) (string, error)            { return s.Pipeline.Track(id) }
func (s *Service) Health() error {
	if s.Store == nil {
		return fmt.Errorf("store unavailable")
	}
	return nil
}
