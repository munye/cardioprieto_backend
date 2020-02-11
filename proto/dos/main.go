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
	campos []Componente `json:campos`
}

// NewEstudio creates a new composite.
func NewEstudio() *Estudio {
	return &Estudio{make([]Componente, 0)}
}

// Add adds a new component to the composite.
func (e *Estudio) Add(componente Componente) {
	e.campos = append(e.campos, componente)
}

// Traverse traverses the composites campos.
func (e *Estudio) Traverse() {
	for i := 0; i < len(e.campos); i++ {
		e.campos[i].Traverse()
	}
}

func main() {

	estudio := NewEstudio()

	for i := 10; i < 100; i++ {
		estudio.Add(NewCampo(i))
	}

	estudio.Traverse()
}
