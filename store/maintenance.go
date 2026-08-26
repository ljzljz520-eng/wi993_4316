package store

import (
	"coffeeware/model"
	"fmt"
	"go.etcd.io/bbolt"
	"time"
)

func (s *Store) PutAll(r model.Record, p model.Profile, e model.Event, a model.Audit) error {
	if err := s.PutRecord(r); err != nil {
		return err
	}
	if err := s.PutProfile(p); err != nil {
		return err
	}
	if err := s.PutEvent(e); err != nil {
		return err
	}
	return s.PutAudit(a)
}
func (s *Store) TouchRecord(id string) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	r.UpdatedAt = time.Now().UTC()
	return s.PutRecord(r)
}
func (s *Store) DeleteRecord(id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets[0]).Delete([]byte(id)) })
}
func (s *Store) Count(bucket string) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n := 0
	if s.db == nil {
		return 0, fmt.Errorf("store closed")
	}
	e := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket")
		}
		return b.ForEach(func(_, v []byte) error {
			if v != nil {
				n++
			}
			return nil
		})
	})
	return n, e
}
func (s *Store) Snapshot() ([]model.Record, error) { return s.ListRecords() }
