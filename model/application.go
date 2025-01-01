// Package model holds all entities which are needed for generation of the
// target application
package model

/*
import (
	. "echoscuffold/model"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"text/template"
	""
	. "github.com/archeopternix/filegenerator"
	"github.com/gertd/go-pluralize"

)*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

// Application holds the global configuration and settings besides all entitites
// for template generation.
// The initialisation and subsequent LoadFromFile() calls creates the Application
// itself and loads the settings and configuration from a YAML file.
//
type Application struct {
	Name   string
	Config struct {
		BasePath string // Basepath in filesystem

	}
	// Settings is the definition of the global attributes
	Settings struct {
		CurrencySymbol    string `yaml:"currency_symbol"`
		DecimalSeparator  string `yaml:"decimal_separator"`
		ThousendSeparator string `yaml:"thousend_separator"`
		TimeFormat        string `yaml:"time_format"`
		DateFormat        string `yaml:"date_format"`
	}
	Entities  map[string]Entity
	Relations map[string]Relation
	Lookups   map[string]LookupList `yaml:"lookups"`
}

type LookupList struct {
	Name string   `yaml:"name"`
	List []string `yaml:"list"`
}

var once sync.Once
var application *Application

// NewApplication creates a new singleton Application or returns the pointer to an existing one
func NewApplication() *Application {
	// call the creation exactly one
	once.Do(func() {
		application = new(Application)
		application.Entities = make(map[string]Entity)
		application.Relations = make(map[string]Relation)
		application.Lookups = make(map[string]LookupList)
	})

	return application
}

// TimeStamp neede for file generation. Will be added in the header of each file
// to track the creation date and time of each file
func (a *Application) TimeStamp() string {
	return time.Now().Format(a.Settings.DateFormat + " " + a.Settings.TimeFormat)
}

// LoadFromFile loads the Application definition from a YAML file and parses all
// dependencies like lookups and master child relationships and therefore creates
// necessary additional entities (e.g. lookup entities) or add additional fields
// (e.g. Id field for every entity)
func (a *Application) LoadFromFile(fname string) error {
	yamlFile, err := ioutil.ReadFile(fname)
	if err != nil {
		return fmt.Errorf("YAML file %v could not be loaded: #%v ", fname, err)
	}
	err = yaml.Unmarshal(yamlFile, a)
	if err != nil {
		return fmt.Errorf("YAML file %v could not be unmarshalled: #%v ", fname, err)
	}

	err = a.parseDependencies()
	return err
}

// StringYAML provides a YAML representation of the pointer to a struct 'a' as string
func StringYAML(a interface{}) string {
	data, err := yaml.Marshal(a)
	if err != nil {
		log.Printf("conversion to YAML string: %v\n", err)
	}
	return string(data)
}

// parseDependencies parse all entities for lookup fields, add unique ID field
// and parse relations between entities and therefore adds dedicated fields for
// parent/child relations and scans for lookups and parent-child relationships
// and therefore creates necessary additional entities (e.g. lookup entities)
// or add additional fields (e.g. Id field for every entity)
func (a *Application) parseDependencies() error {
	for key, entity := range a.Entities {
		for _, field := range entity.Fields {

			// If a lookup field is present check for lookup table
			if field.Kind == "Lookup" {
				if _, ok := a.Lookups[strings.ToLower(field.Lookup)]; !ok {
					return fmt.Errorf("Lookup with name '%s' could not be found", field.Lookup)
				}
				lk := field
				lk.Object = field.Lookup
				a.Entities[key].Fields[strings.ToLower(field.Name)] = lk
			}
		}
	}

	// Add an ID field to all entities
	for key, entity := range a.Entities {
		// add ID field
		entity.Fields["ID"] = Field{Name: "ID", Kind: "Integer", Required: true}
		a.Entities[key] = entity
	}

	// add fields for relationships between entities
	for _, relation := range a.Relations {
		if relation.Kind == "One-to-Many" {
			// add child field
			childentity := a.Entities[strings.ToLower(relation.Child)]
			childentity.Fields[relation.Parent+"ID"] = Field{Name: relation.Parent + "ID", Kind: "Child", Object: relation.Parent}
			a.Entities[strings.ToLower(relation.Child)] = childentity
			// add parent field
			parententity := a.Entities[strings.ToLower(relation.Parent)]
			parententity.Fields[relation.Child] = Field{Name: relation.Child, Kind: "Parent", Object: relation.Child}
			a.Entities[strings.ToLower(relation.Parent)] = parententity
		}
	}

	/*	for _,lookup := range a.Lookups {
			// do nothing so far
		}
	*/

	return nil
}
