// golangcrud project main.go
package main

import (
	"fmt"
	. "golangcrud/model"
)

func main() {
	p := NewProject("Erstes Projekt")
	Projects[p.Name] = *p

	fmt.Println(*p)
}
