package model

import "time"

type Record struct {
	ID, Name, Category, Material, Status string
	Price                                float64
	Stock                                int
	CreatedAt, UpdatedAt                 time.Time
}
type Profile struct {
	ID, DisplayName, Email, Role string
	Active                       bool
	CreatedAt                    time.Time
}
type Event struct {
	ID, RecordID, Kind, Actor, Details string
	At                                 time.Time
}
type Audit struct {
	ID, Action, Subject, Result, Message string
	At                                   time.Time
}

func NewRecord(id, name, category string, price float64, stock int) Record {
	now := time.Now().UTC()
	return Record{ID: id, Name: name, Category: category, Price: price, Stock: stock, Status: "draft", CreatedAt: now, UpdatedAt: now}
}
func (r Record) Valid() error {
	if r.ID == "" || r.Name == "" || r.Category == "" {
		return ErrInvalidRecord
	}
	if r.Price < 0 || r.Stock < 0 {
		return ErrInvalidRecord
	}
	return nil
}
func (r *Record) Publish() error {
	if err := r.Valid(); err != nil {
		return err
	}
	if r.Stock == 0 {
		return ErrOutOfStock
	}
	r.Status = "published"
	r.UpdatedAt = time.Now().UTC()
	return nil
}
func (r *Record) Archive() error {
	if r.Status == "archived" {
		return ErrAlreadyArchived
	}
	r.Status = "archived"
	r.UpdatedAt = time.Now().UTC()
	return nil
}
func (p Profile) Can(action string) bool {
	if !p.Active {
		return false
	}
	if p.Role == "admin" || p.Role == "editor" {
		return action != "delete" || p.Role == "admin"
	}
	return action == "read"
}
func NewEvent(id, recordID, kind, actor, details string) Event {
	return Event{ID: id, RecordID: recordID, Kind: kind, Actor: actor, Details: details, At: time.Now().UTC()}
}
func NewAudit(id, action, subject, result, msg string) Audit {
	return Audit{ID: id, Action: action, Subject: subject, Result: result, Message: msg, At: time.Now().UTC()}
}
