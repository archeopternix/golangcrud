// GenerationDSL project main.go
package main

import (
	//	"fmt"
	. "golangcrud/model"
	"log"
)

func main() {
	a := NewApplication()

	if err := a.LoadFromFile("projectmgnt.yaml"); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	//	fmt.Println(a.StringYAML())

	c := NewGenerator()
	/*	if err := c.AddModule("modules/application/app.yaml"); err != nil {
			log.Fatalf("ERROR: %v", err)
		}
		if err := c.AddModule("modules/model/models.yaml"); err != nil {
			log.Fatalf("ERROR: %v", err)
		}

		 if err := c.AddModule("modules/mockdatabase/mockdatabase.yaml"); err != nil {
			log.Fatalf("ERROR: %v", err)
		}
			if err := c.AddModule("modules/database/database.yaml"); err != nil {
				log.Fatalf("ERROR: %v", err)
			}

				if err := c.AddModule("modules/view/view.yaml"); err != nil {
			log.Fatalf("ERROR: %v", err)
		}
	*/

	if err := c.AddModule("modules/databasetest/databasetest.yaml"); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	if err := c.AddModule("modules/mockdatabase/mockdatabase.yaml"); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	if err := c.GenerateAll(a); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	//	fmt.Println(c.StringYAML())

}
