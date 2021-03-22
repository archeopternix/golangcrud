{{define "mockrepo" -}}
// Package mockdatabase contains structures and function for mock database access
// Generated code - do not modify it will be overwritten!!
// Time: {{.Entity.TimeStamp}}
package mockdatabase

import (
	"fmt"
	. "{{.AppName}}/model"
)

{{with .Entity}}

// {{.Name}}Repo is the interface for a {{.Name}} repository that will persist 
// and retrieve data and has to be implemented for concrete Databases 
// (e.g. db *sqlx.DB) or other respositories
type {{.Name}}Repo struct{
	data map[uint64]{{.Name}}
	count uint64
}

var {{.Name | lowercase}}repo *{{.Name}}Repo


func New{{.Name}}Repo() *{{.Name}}Repo {
	{{.Name | lowercase}}repo = new({{.Name}}Repo)
	{{.Name | lowercase}}repo.data = make(map[uint64]{{.Name}})
	{{.Name | lowercase}}repo.count = 1
	return {{.Name | lowercase}}repo
}


// Get queries a {{.Name | lowercase}} by id, throws an error when id is not found
func (repo {{.Name}}Repo) Get(id uint64) (*{{.Name}}, error) {
	value, ok := {{.Name | lowercase}}repo.data[id]
	if !ok {
		return nil, fmt.Errorf("get project with id %d, record not found", id)
	}
	return &value, nil
}

// GetAll returns all records ordered by the fields  with isLabel=true
func (repo {{.Name}}Repo) GetAll() ({{.Name}}List, error) {
	var list {{.Name}}List
	for _,value:=range {{.Name | lowercase}}repo.data {
		{{- range .Fields}}{{if eq .Kind "Lookup" "Child"}}
		if {{.Object | lowercase}},_:= {{.Object | lowercase}}repo.Get(value.{{.Name}}); {{.Object | lowercase}}!=nil {
			value.{{.Object}} = {{.Object | lowercase}}.Label()
		}
		{{- end}}{{end}}
		list = append(list,value)
	}
			
	return list, nil
}

// Delete deletes the {{.Name | lowercase}} with id, throws an error when id is not found
func (repo {{.Name}}Repo) Delete(id uint64) error {
	_, ok := {{.Name | lowercase}}repo.data[id]
	if !ok {
		return fmt.Errorf("delete project with id '%d', record not found", id)
	}

	delete({{.Name | lowercase}}repo.data, id)
	return nil
}

// Update updates all fields in the database table with data from *{{.Name}})
func (repo {{.Name}}Repo) Update({{.Name | lowercase}} *{{.Name}}) error {
	_, ok := {{.Name | lowercase}}repo.data[{{.Name | lowercase}}.ID]
	if !ok {
		return fmt.Errorf("update project with id '%d', record not found", {{.Name | lowercase}}.ID)
	}	
	{{.Name | lowercase}}repo.data[{{.Name | lowercase}}.ID] = *{{.Name | lowercase}}
	return nil
}

// Insert inserts a new record in the database table with data from *{{.Name}})
func (repo {{.Name}}Repo) Insert({{.Name | lowercase}} *{{.Name}}) error {
	{{.Name | lowercase}}repo.count +=1
	{{.Name | lowercase}}.ID ={{.Name | lowercase}}repo.count
	{{.Name | lowercase}}repo.data[{{.Name | lowercase}}repo.count] = *{{.Name | lowercase}}
	return nil
}

// GetLabels returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{.Name}}Repo) GetLabels() (Labels, error) {
	labels := make(Labels)
	for _, value := range {{.Name | lowercase}}repo.data {
		labels[value.ID] = value.Label()
	}
	return labels, nil
}

{{$name:=.Name}}
{{- range .Fields}}{{if eq .Kind "Parent"}}
// GetAll{{$name | plural}}ForParentID returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{$name}}Repo) GetAll{{.Name | plural}}ByParentID(parentID uint64) ({{.Name}}List)	{
	list := {{.Name}}List{}
	{{.Name | plural | lowercase}}, err := {{.Name | lowercase}}repo.GetAll()
	if err!=nil {
		return list
	}
	for _, {{.Name | lowercase}} := range {{.Name | plural | lowercase}} {
		if {{.Name | lowercase}}.{{$name}}ID == parentID {
			list = append(list, {{.Name | lowercase}})
		}
	}
	return list
}			
{{- end}}{{end}}

{{end}}
{{end}}


