{{define "mockrepo" -}}
// Package database contains structures and function for generic database access
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
	data {{.Name}}List
}

var {{.Name | lowercase}}Repo *{{.Name}}Repo

func New{{.Name}}Repo() *{{.Name}}Repo {
	{{.Name | lowercase}}Repo = new({{.Name}}Repo)
	return {{.Name | lowercase}}Repo
}


// Get queries a {{.Name | lowercase}} by id, throws an error when id is not found
func (repo {{.Name}}Repo) Get(id uint64) (*{{.Name}}, error) {
	if id > uint64(len(repo.data)) {
		return nil, fmt.Errorf("get {{.Name | lowercase}} with id %d, only %d count records", id, len(repo.data))	
	}
	return &repo.data[id], nil
}

// GetAll returns all records ordered by the fields  with isLabel=true
func (repo {{.Name}}Repo) GetAll() ({{.Name}}List, error) {
	return repo.data, nil
}

// Delete deletes the {{.Name | lowercase}} with id, throws an error when id is not found
func (repo {{.Name}}Repo) Delete(id uint64) error {
	if id > uint64(len(repo.data)) {
		return fmt.Errorf("get {{.Name | lowercase}} with id %d, only %d count records", id, len(repo.data))	
	}
	return nil
}

// Update updates all fields in the database table with data from *{{.Name}})
func (repo {{.Name}}Repo) Update({{.Name | lowercase}} *{{.Name}}) error {
	repo.data[{{.Name | lowercase}}.ID]= *{{.Name | lowercase}}
	return nil
}

// Insert inserts a new record in the database table with data from *{{.Name}})
func (repo {{.Name}}Repo) Insert({{.Name | lowercase}} *{{.Name}}) error {
	{{.Name | lowercase}}.ID = uint64(len(repo.data))
	repo.data = append(repo.data,*{{.Name | lowercase}})
	return nil
}

// GetLabelsFor returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{.Name}}Repo) GetLabels() (Labels, error) {
	l := make(Labels)
	l[1]="Alpha"
	l[2]="Beta"
	l[3]="Gamma"
	return l, nil
}

{{$name:=.Name}}
{{- range .Fields}}{{if eq .Kind "Parent"}}
// GetAll{{$name | plural}}ForParentID returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{$name}}Repo) GetAll{{.Name | plural}}ByParentID(parentID uint64) ({{.Name}}List, error)	{
	return {{.Name | lowercase}}Repo.GetAll()
}			
{{- end}}{{end}}

{{end}}
{{end}}


