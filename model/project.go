//Package model project.go holds all struct and function to the project
package model

// Project relates to an 'Object' or struct
type Project struct {
	Name     string            `json:"name"` // unique identifier for project name
	Entities map[string]Entity `json:"entities"`
}

// Projects holds all projects in a map. Projects can be added by 'AddProject'
var Projects map[string]Project

// NewProject creates a new Project and initializes the Entity map
func NewProject(name string) *Project {
	project := new(Project)
	project.Name = name
	project.Entities = make(map[string]Entity)
	return project
}

// ErrEntryAlreadyExists customer error type to indicate when an entry already exists
type ErrEntryAlreadyExists struct {
	name string
}

// NewErrEntryAlreadyExists creates oa new error of type ErrEntryAlreadyExists
func NewErrEntryAlreadyExists(name string) *ErrEntryAlreadyExists {
	return &ErrEntryAlreadyExists{
		name: name,
	}
}

func (e *ErrEntryAlreadyExists) Error() string {
	return "the entry already exists: " + e.name
}

// ProjectExists checks if a project with 'name' exists in the global 'Project' map
func ProjectExists(name string) bool {
	_, ok := Projects[name]
	if ok {
		return true
	}
	return false
}

// AddProject adds a project to the global 'Project' map and checks
func AddProject(p *Project) error {
	if ProjectExists(p.Name) {
		return NewErrEntryAlreadyExists(p.Name)
	}
	Projects[p.Name] = *p
	return nil
}

func init() {
	Projects = make(map[string]Project)
}
