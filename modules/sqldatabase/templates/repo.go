{{define "repo" -}}
// Package database contains structures and function for generic database access
// Generated code - do not modify it will be overwritten!!
// Time: {{.Entity.TimeStamp}}
package database

import (
	"github.com/jmoiron/sqlx"	
	"fmt"
	model "{{.AppName}}/model"
)

{{with .Entity}}

// {{.Name}}Repo is the interface for a {{.Name}} repository that will persist 
// and retrieve data and has to be implemented for concrete Databases 
// (e.g. db *sqlx.DB) or other respositories
type {{.Name}}Repo struct{
	// pointer to the global database
	DB *sqlx.DB
}

// Get queries a {{.Name | lowercase}} by id, throws an error when id is not found
func (repo {{.Name}}Repo) Get(id uint64) (*model.{{.Name}}, error) {
	{{.Name | lowercase}} := new(model.{{.Name}})
	if err := db.Get({{.Name | lowercase}}, "SELECT * FROM {{.Name | lowercase | plural}} WHERE id=$1", id); err != nil {
		return nil, fmt.Errorf("get {{.Name | lowercase}} with id %d, %v", id, err)
	}
	return {{.Name | lowercase}}, nil
}

// GetAll returns all records ordered by the fields  with isLabel=true
func (repo {{.Name}}Repo) GetAll() (model.{{.Name}}List, error) {
	list := make(model.{{.Name}}List)

	rows, err := db.Queryx("SELECT * FROM {{.Name | lowercase| plural}} ORDER BY {{range $index, $field:=.Fields}}{{if eq .IsLabel true}}{{if gt $index 0}},{{end}}{{$field.Name}} {{end}}{{end}} ASC")
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		{{.Name | lowercase}} := new(model.{{.Name}})
		if err := rows.StructScan({{.Name | lowercase}}); err != nil {
			return nil, fmt.Errorf("parsing {{.Name | lowercase| plural}} struct, err %v", err)
		}
		list = append(list, {{.Name | lowercase}})
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	return list, nil
}

// Delete deletes the {{.Name | lowercase}} with id, throws an error when id is not found
func (repo {{.Name}}Repo) Delete(id uint64) error {
	deleteStatement := fmt.Sprintf("DELETE FROM {{.Name | lowercase | plural}} WHERE id=%d", id)
	if _, err := db.Exec(deleteStatement); err != nil {
		return fmt.Errorf("delete {{.Name | lowercase}} with id %d, %v", id, err)
	}
	return nil
}

// Update updates all fields in the database table with data from *{{.Name}})
func (repo {{.Name}}Repo) Update({{.Name | lowercase}} *model.{{.Name}}) error {
	updateStatement := {{template "repoupdate" .}}
	if _, err := db.NamedExec(updateStatement, {{.Name | lowercase}}); err != nil {
		return fmt.Errorf("update {{.Name | lowercase| plural}}, %v", err)
	}
	return nil
}

// Insert inserts a new record in the database table with data from *{{.Name}})
func (repo {{.Name}}Repo) Insert({{.Name | lowercase}} *model.{{.Name}}) error {
	insertStatement := {{template "repoinsert" .}}
	if _, err := db.NamedExec(insertStatement, {{.Name | lowercase}}); err != nil {
		return fmt.Errorf("insert {{.Name | lowercase| plural}}, %v", err)
	}
	return nil
}

// GetLabelsFor returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{.Name}}Repo) GetLabels() (model.Labels, error) {
	l := make(model.Labels)

	rows, err := db.Queryx("SELECT * FROM {{.Name | lowercase| plural}} ORDER BY {{range $index, $field:=.Fields}}{{if eq .IsLabel true}}{{if gt $index 0}},{{end}}{{$field.Name}} {{end}}{{end}} ASC")
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		{{.Name | lowercase}} := new(model.{{.Name}})
		if err := rows.StructScan({{.Name | lowercase}}); err != nil {
			return nil, fmt.Errorf("parsing {{.Name | lowercase| plural}} struct, err %v", err)
		}
		{{$name:= .Name}}
		label := fmt.Sprintf("{{range .Fields}}{{if eq .IsLabel true}}%s {{end}}{{end}}"{{range .Fields}}{{if eq .IsLabel true}},{{$name | lowercase}}.{{.Name}} {{end}}{{end}})
		l[{{.Name | lowercase}}.ID] = label
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	return l, nil
}

{{$name:=.Name}}
{{- range .Fields}}{{if eq .Kind "Parent"}}
// GetAll{{$name | plural}}ForParentID returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{$name}}Repo) GetAll{{.Name | plural}}ByParentID(parentID uint64) (model.{{.Name}}List, error)	{
	list := make(model.{{.Name}}List)

	query:= fmt.Sprintf("SELECT * FROM {{.Name | lowercase| plural}} WHERE id=%d", parentID)
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("selecting {{.Name | lowercase| plural}} with id: '%d' from database, %v",parentID, err)
	}
	
	for rows.Next() {
		{{.Name | lowercase}} := new(model.{{.Name}})
		if err := rows.StructScan({{.Name | lowercase}}); err != nil {
			return nil, fmt.Errorf("parsing {{.Name | lowercase| plural}} struct, %v", err)
		}
		list = append(list, {{.Name | lowercase}})
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	return list, nil
}			
{{- end}}{{end}}

{{end}}
{{end}}

