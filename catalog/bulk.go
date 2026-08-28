package catalog

import (
	"coffeeware/model"
	"fmt"
)

type ImportResult struct {
	Accepted, Rejected int
	Errors             []string
}

func (c *Catalog) Import(rows []model.Record) ImportResult {
	out := ImportResult{Errors: []string{}}
	for i, r := range rows {
		if e := c.Register(r); e != nil {
			out.Rejected++
			out.Errors = append(out.Errors, fmt.Sprintf("row %d: %v", i, e))
		} else {
			out.Accepted++
		}
	}
	return out
}
func (c *Catalog) PublishMany(ids []string) (int, []error) {
	ok := 0
	errs := []error{}
	for _, id := range ids {
		if _, e := c.Publish(id); e != nil {
			errs = append(errs, e)
		} else {
			ok++
		}
	}
	return ok, errs
}
func (c *Catalog) ArchiveMany(ids []string) (int, []error) {
	ok := 0
	errs := []error{}
	for _, id := range ids {
		if _, e := c.Archive(id); e != nil {
			errs = append(errs, e)
		} else {
			ok++
		}
	}
	return ok, errs
}
func (c *Catalog) RestockMany(items map[string]int) (int, []error) {
	ok := 0
	errs := []error{}
	for id, q := range items {
		if _, e := c.Restock(id, q); e != nil {
			errs = append(errs, e)
		} else {
			ok++
		}
	}
	return ok, errs
}
