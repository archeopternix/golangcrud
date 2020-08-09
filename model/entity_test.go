package model

import (
	"testing"
)

func TestNewEntity(t *testing.T) {
	err := AddProject(NewProject("FieldProject"))
	if err != nil {
		t.Errorf(err.Error())
	}
	project, ok := Projects["FieldProject"]
	if !ok {
		t.Errorf("the project could not added and retrieved: %v", Projects)
		return
	}

	entity := NewEntity()
	entity.Name = "Oger"
	project.Entities["Oger"] = *entity
	_, hit := project.Entities["Oger"]
	if !hit {
		t.Errorf("the entity could not be found: %v", Projects)
		return
	}

	t.Logf("entity successful added to a project")
}

func TestNewField(t *testing.T) {
	project, ok := Projects["FieldProject"]
	if !ok {
		t.Errorf("the project could not added and retrieved: %v", Projects)
		return
	}

	entity, hit := project.Entities["Oger"]
	if !hit {
		t.Errorf("the entity could not be found: %v", Projects)
		return
	}

	err := entity.AddField(Field{Name: "one"})
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("field successful added to an entity")
}
