package catalog

import (
	"coffeeware/model"
	"coffeeware/store"
	"sort"
	"strings"
	"time"
)

type Catalog struct{ st *store.Store }

func New(st *store.Store) *Catalog { return &Catalog{st: st} }
func (c *Catalog) Register(r model.Record) error {
	if e := r.Valid(); e != nil {
		return e
	}
	return c.st.PutRecord(r)
}
func (c *Catalog) Find(id string) (model.Record, error) { return c.st.GetRecord(id) }
func (c *Catalog) Search(term, category, status string) ([]model.Record, error) {
	all, e := c.st.ListRecords()
	if e != nil {
		return nil, e
	}
	out := make([]model.Record, 0)
	term = strings.ToLower(term)
	for _, r := range all {
		if term != "" && !strings.Contains(strings.ToLower(r.Name), term) && !strings.Contains(strings.ToLower(r.Material), term) {
			continue
		}
		if category != "" && r.Category != category {
			continue
		}
		if status != "" && r.Status != status {
			continue
		}
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.After(out[j].UpdatedAt) })
	return out, nil
}
func (c *Catalog) Publish(id string) (model.Record, error) {
	r, e := c.Find(id)
	if e != nil {
		return r, e
	}
	if e = r.Publish(); e != nil {
		return r, e
	}
	return r, c.st.PutRecord(r)
}
func (c *Catalog) Archive(id string) (model.Record, error) {
	r, e := c.Find(id)
	if e != nil {
		return r, e
	}
	if e = r.Archive(); e != nil {
		return r, e
	}
	return r, c.st.PutRecord(r)
}
func (c *Catalog) Restock(id string, qty int) (model.Record, error) {
	r, e := c.Find(id)
	if e != nil {
		return r, e
	}
	if qty <= 0 {
		return r, model.ErrInvalidRecord
	}
	r.Stock += qty
	r.UpdatedAt = now()
	return r, c.st.PutRecord(r)
}
func now() time.Time { return time.Now().UTC() }
