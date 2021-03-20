{{define "mockrepotest" -}}
// Package mockdatabase contains structures and function for mock database access
// Generated code - do not modify it will be overwritten!!
// Time: {{.Entity.TimeStamp}}
package mockdatabase

import (
	"testing"
	"math/rand"

	model "{{.AppName}}/model"
)

{{with .Entity}}


func Test{{.Name}}RepoPositive(t *testing.T) {
	record := &model.{{.Name}}{{"{"}} {{range .Fields}}{{template "mockrepotesttypes" .}}{{end}}{{"}"}}
	if err := {{.Name | lowercase}}db.Insert(record); err != nil {
		t.Errorf("adding {{.Name}} failed %v, %d", err, rand.Int())
	}
	list, err := {{.Name | lowercase}}db.GetAll()
	if err != nil {
		t.Errorf("cannot get records from db %v", err)
	}
	if len(list) != 1 {
		t.Errorf("count records should be 1 is %d", len(list))
	} else {
		t.Logf("successfully inserted and retrieved 1 record from db")
	}

	record, err = {{.Name | lowercase}}db.Get(list[0].ID)
	if err != nil {
		t.Errorf("cannot get record %d from db %v", list[0].ID, err)
	} else {
		t.Logf("successfully retrieved record %d from db", record.ID)
	}

	err = {{.Name | lowercase}}db.Update(record)
	if err != nil {
		t.Errorf("cannot update %d from db %v", record.ID, err)
	} else {
		t.Logf("successfully updated record %d from db", record.ID)
	}

	err = {{.Name | lowercase}}db.Delete(record.ID)
	if err != nil {
		t.Errorf("cannot delete %d from db %v", record.ID, err)
	}

	list, err = {{.Name | lowercase}}db.GetAll()
	if err != nil {
		t.Errorf("cannot get records from db %v", err)
	}
	if len(list) > 0 {
		t.Errorf("no records expected, count :%d", len(list))
	} else {
		t.Logf("successfully deleted record %d from db", record.ID)
	}
}

func Test{{.Name}}RepoNegative(t *testing.T) {
	list, err := {{.Name | lowercase}}db.GetAll()
	if err != nil {
		t.Errorf("cannot get records from db %v", err)
	}
	if len(list) > 1 {
		t.Errorf("count records should be 1 is %d", len(list))
	} else {
		t.Logf("as expected no records found")
	}

	_, err = {{.Name | lowercase}}db.Get(65534)
	if err == nil {
		t.Errorf("expected not to find a record")
	} else {
		t.Logf("as expected record not found and error returned")
	}

	record := &model.{{.Name}}{{"{"}} {{range .Fields}}{{template "mockrepotesttypes" .}}{{end}}{{"}"}}
	err = {{.Name | lowercase}}db.Update(record)
	if err == nil {
		t.Errorf("expected not to update a record and throw an error")
	} else {
		t.Logf("as expected no record updated and error returned")
	}

	err = {{.Name | lowercase}}db.Delete(record.ID)
	if err == nil {
		t.Errorf("expected not to delete a record and throw an error")
	} else {
		t.Logf("as expected no record deleted and error returned")
	}
}
{{end}}
{{end}}





