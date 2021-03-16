// GenerationDSL project main.go
package main

import (
	"fmt"
	. "golangcrud/model"
	"log"
)

func main() {
	a := NewApplication()
	if err := a.LoadFromFile("projectmgnt.yaml"); err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(a.StringYAML())

	c := NewGenerator()
	if err := c.AddModule("modules/model/models.yaml"); err != nil {
		fmt.Println(err)
	}

	if err := c.AddModule("modules/database/database.yaml"); err != nil {
		fmt.Println(err)
	}

	if err := c.GenerateAll(a); err != nil {
		fmt.Println(err)
	}

	//	fmt.Println(c.StringYAML())

}
