package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Estudio describes the behavior that needs to be exercised uniformly
// across all primitive and composite objects.
type Componente interface {
	Traverse()
}

// Campo describes a primitive leaf object in the hierarchy.
type Campo struct {
	Value string `json:valor`
}

// NewCampo creates a new leaf.
func NewCampo(value string) *Campo {
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
	PacienteID uint         `json:paciente_id`
	Campos     []Componente `json:campos`
}

// NewEstudio creates a new composite.
func NewEstudio(pacienteID uint) *Estudio {
	return &Estudio{PacienteID: pacienteID, Campos: make([]Componente, 0)}
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

	estudio := NewEstudio(1)

	for i := 0; i < 10; i++ {
		estudio.Add(NewCampo(strconv.Itoa(i)))
	}

	estudio.Traverse()

}
