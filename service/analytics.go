package service

import (
	"coffeeware/catalog"
	"coffeeware/model"
	"sort"
)

func (s *Service) Metrics() (catalog.Metrics, error) {
	rs, e := s.Store.ListRecords()
	if e != nil {
		return catalog.Metrics{}, e
	}
	return catalog.Summarize(rs), nil
}
func (s *Service) TopCategories() ([]string, error) {
	rs, e := s.Store.ListRecords()
	if e != nil {
		return nil, e
	}
	m := map[string]int{}
	for _, r := range rs {
		m[r.Category]++
	}
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out, nil
}
func (s *Service) ValidateBatch(rs []model.Record) []error {
	errs := []error{}
	for _, r := range rs {
		if e := r.Valid(); e != nil {
			errs = append(errs, e)
		}
	}
	return errs
}
func (s *Service) Audit(action, subject, result, msg string) error {
	return s.Store.PutAudit(model.NewAudit(subject, action, subject, result, msg))
}
