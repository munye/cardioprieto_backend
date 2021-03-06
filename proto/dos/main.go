package main

import (
	"encoding/json"
	"fmt"
)

// Estudio describes the behavior that needs to be exercised uniformly
// across all primitive and composite objects.
type Componente interface {
	Traverse()
}

// Campo describes a primitive leaf object in the hierarchy.
type Campo struct {
	Value int `json:valor`
}

// NewCampo creates a new leaf.
func NewCampo(value int) *Campo {
	return &Campo{Value: value}
}

// Traverse prints the value of the leaf.
func (c *Campo) Traverse() {
	js, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(js))
}

// Estudio describes a composite of components.
type Estudio struct {
	Title  string       `json:title`
	Campos []Componente `json:campos`
}

// NewEstudio creates a new composite.
func NewEstudio() *Estudio {
	return &Estudio{Title: "turulu", Campos: make([]Componente, 0)}
}

// Add adds a new component to the composite.
func (e *Estudio) Add(componente Componente) {
	e.Campos = append(e.Campos, componente)
}

// Traverse traverses the composites campos.
func (e *Estudio) Traverse() {
	js, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))
}

func main() {

	estudio := NewEstudio()

	for i := 0; i < 10; i++ {
		estudio.Add(NewCampo(i))
	}

	estudio.Traverse()

}
