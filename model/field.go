//Package model
package model

// Field is each and every single attribute.
// Object is empty except in case type=lookup or child keeps the name of the Object
type Field struct {
	Name      string `json:"name"`
	Type      string `json:"type"`             // string, int, bool, lookup, tel, email
	Object    string `json:"object,omitempty"` // for lookup, child relations - mappingtable for many2many relations
	Maxlength int    `json:"maxlength,omitempty"`
	Size      int    `json:"size,omitempty"` // for textarea size = cols
	Required  bool   `json:"required"`
	Step      int    `json:"step,omitempty"`    //for Number fields
	Min       int    `json:"min,omitempty"`     //for Number fields
	Max       int    `json:"max,omitempty"`     //for Number fields
	Rows      int    `json:"rows,omitempty"`    //for textarea
	IsLabel   bool   `json:"islabel,omitempty"` // when true is the shown text for select boxes
}
