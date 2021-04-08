# CRUD Generator: golangcrud 
Application that generates a CRUD application with selectable modules for backend (repositories) and frontend (REST or web based application)

## Installation

	go get github.com/archeopternix/golangcrud

## Configuration File
The configuration is done by writing a YAML File

### General Information
Basic configuration settings for applications like name and the basepath where the target application should be generated to

    config:
	name: SampleApp
	basepath: \Users\A.Eisner\go\src
	settings:
	  decimal_separator: ","
	  currency_symbol: '€'
	  thousend_separator: "."
	  time_format: 15:04:05.000
	  date_format: 02.01.2006

### Simple Entity
Simple entity with the name and a list of field definitions. An ID field of type uint64 will be generated automatically which serves as primary + unique field in the database

	user:
	  name: User
	  kind: default
	  fields:
	  - name: Lastname
	    kind: Text
	    length: 45
	    required: true
	    islabel: true
	  - name: Username
	    kind: Text
	    length: 45
	    required: true

### Entity with Lookup
Entity with a field of kind 'Lookup'. Lookup fields will trigger the creation of an automatic generated lookup entity. 
An ID field of type uint64 will be generated automatically which serves as primary + unique field in the database

	task:
	  name: Task
	  kind: default
	  fields:
	  - name: Description
	    kind: Longtext
	    length: 200
	    required: true
	    islabel: true
	  - name: Tasktype
	    kind: Lookup

### Relationships
Relationship between different Entitites (currently supported 1..n). 
1..n: for this kind of relationship there will be generated fields on the one and many side automatically. 
Parent side (..n) the will be a slice created that holds all mapped entries. 
Child side (1..) the will be added a ParentID (in the case below ProjectID) field that holds the ID of the parent entity

	relations:
	- kind: one_to_many
	  parent: Project
	  child: Task
	- kind: one_to_many
	  parent: Project
	  child: User

## Modules

### [Application](modules/application/README.md)

### [Model](modules/model/README.md)

### View based on echo and Bulma.css

### Database Test

### Mockdatabase

### SQL Database

