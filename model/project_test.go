package model

import (
	"testing"
)

func TestNewProject(t *testing.T) {
	p := NewProject("TestProject")
	err := AddProject(p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	_, ok := Projects["TestProject"]
	if !ok {
		t.Errorf("the project could not added and retrieved %v", Projects)
		return
	}
	t.Logf("project successful added and retrieved.")
}

func TestDuplicateAddProject(t *testing.T) {
	p := NewProject("DuplicateProject")
	err := AddProject(p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	err = AddProject(p)
	if err == nil {
		t.Errorf("duplicate entry not catched")
		return
	}
	t.Logf("success that duplicate entry rejected")
}

func TestProjectExists(t *testing.T) {
	if ProjectExists("ProjectOne") {
		t.Errorf("project must not yet exist: %v", Projects)
		return
	}
	t.Logf("as expected project does not yet exist")

	err := AddProject(NewProject("ProjectOne"))
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if !ProjectExists("ProjectOne") {
		t.Errorf("the project could not added and retrieved %v", Projects)
		return
	}

	t.Logf("success project added and existing.")
}
