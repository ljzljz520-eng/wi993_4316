package store

import (
	"coffeeware/model"
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
)

type Transaction struct{ tx *bbolt.Tx }

func (s *Store) WithTransaction(fn func(*Transaction) error) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return fn(&Transaction{tx: tx}) })
}
func (t *Transaction) Record(r model.Record) error {
	b := t.tx.Bucket(buckets[0])
	v, e := json.Marshal(r)
	if e != nil {
		return e
	}
	return b.Put([]byte(r.ID), v)
}
func (t *Transaction) Event(e model.Event) error {
	b := t.tx.Bucket(buckets[2])
	v, x := json.Marshal(e)
	if x != nil {
		return x
	}
	return b.Put([]byte(e.ID), v)
}
func (t *Transaction) Audit(a model.Audit) error {
	b := t.tx.Bucket(buckets[3])
	v, x := json.Marshal(a)
	if x != nil {
		return x
	}
	return b.Put([]byte(a.ID), v)
}
func (t *Transaction) Exists(bucket, id string) bool {
	return t.tx.Bucket([]byte(bucket)).Get([]byte(id)) != nil
}
