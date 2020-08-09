// Package model
package model

import (
	"time"
)

// Entity relates to an 'Object' that will have the CRUD functionality and will
// be generated into a struct
type Entity struct {
	Name       string                  `json:"name"` // unique name for an entity
	Fields     map[string]Field        `json:"fields"`
	EntityType `json:"type,omitempty"` // 0..Standard, 1..Lookup
}

//
type EntityType int

const (
	Standard EntityType = iota
	Lookup
)

func (d EntityType) String() string {
	return [...]string{"Standard", "Lookup"}[d]
}

// Relation is the linking between 2 entities and can have different kinds of
// relation types
type Relation struct {
	Id     string `json:"id"`
	Parent string `json:"parent"`
	Child  string `json:"child"`
	Kind   string `json:"kind"` // "one2many", "many2many"
}

func NewEntity() *Entity {
	e := new(Entity)
	e.Fields = make(map[string]Field)
	return e
}

// FieldExists checks if a field with 'name' exists in the entity 'Fields' map
func (e *Entity) FieldExists(name string) bool {
	_, ok := e.Fields[name]
	if ok {
		return true
	}
	return false
}

func (e *Entity) AddField(f Field) error {
	if e.FieldExists(f.Name) {
		return NewErrEntryAlreadyExists(f.Name)
	}
	e.Fields[f.Name] = f
	return nil
}

func (e Entity) TimeStamp() string {
	return time.Now().Format(time.UnixDate)
}
