// Package model holds all entities which are needed for generation of the
// target application
package model

import (
	"time"
)

// Entity relates to an 'Object' or struct
type Entity struct {
	Name   string  `yaml:"name"`
	Fields []Field `yaml:"fields"`
	Kind   string  `yaml:"type,omitempty"` // 0..Normal, 1..Lookup 2..Many2Many
}

// Field is each and every single attribute.
// Object is empty except in case type=lookup or child keeps the name of the Object
type Field struct {
	Name     string `yaml:"name"`
	Kind     string `yaml:"kind"` // string, int, bool, lookup, tel, email
	Required bool   `yaml:"required,omitempty"`
	IsLabel  bool   `yaml:"islabel,omitempty"` // when true is the shown text for select boxes
	Object   string `yaml:"object,omitempty"`  // for lookup, child relations - mappingtable for many2many relations
	Length   int    `yaml:"length,omitempty"`
	Size     int    `yaml:"size,omitempty"` // for textarea size = cols

	Step int `yaml:"step,omitempty"` //for Number fields
	Min  int `yaml:"min,omitempty"`  //for Number fields
	Max  int `yaml:"max,omitempty"`  //for Number fields

	Rows int `yaml:"rows,omitempty"` //for textarea
}

// Relation holds the definition for parent - child relationships.
// When parsed by Application additional fields will be added to the child and parent
// entities
type Relation struct {
	Parent string `json:"parent"`
	Child  string `json:"child"`
	Kind   string `json:"kind"` // "one2many", "many2many"
}

// AddField add fields to an entity
func (e *Entity) AddField(f Field) {
	e.Fields = append(e.Fields, f)
}

// TimeStamp neede for file generation. Will be added in the header of each file
// to track the creation date and time of each file
func (e Entity) TimeStamp() string {
	return time.Now().Format(application.Settings.DateFormat + " " + application.Settings.TimeFormat)
}

/* Testcode:
e := NewEntity()
e.Name = "Role"
e.addField(Field{Name: "Id", Type: inttype, Object: ""})
e.addField(Field{Name: "Name", Type: stringtype, Object: ""})
e.addField(Field{Name: "Accounts", Type: slicetype, Object: "Account"})

err := Database.Insert(e)
if err != nil {
	panic(err)
}

_, es := getAllEntities()
*/
