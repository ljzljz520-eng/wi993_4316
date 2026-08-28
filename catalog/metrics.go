package catalog

import (
	"coffeeware/model"
	"time"
)

type Metrics struct {
	Total, Published, Archived, Stock int
	Updated                           time.Time
}

func Summarize(rs []model.Record) Metrics {
	m := Metrics{Updated: time.Now().UTC()}
	for _, r := range rs {
		m.Total++
		m.Stock += r.Stock
		if r.Status == "published" {
			m.Published++
		}
		if r.Status == "archived" {
			m.Archived++
		}
	}
	return m
}
