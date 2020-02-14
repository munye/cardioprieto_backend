package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	//"github.com/munye/prueba_backend_go/common"
	"strconv"
)

// Estudio describes the behavior that needs to be exercised uniformly
// across all primitive and composite objects.
type Componente interface {
	Traverse()
}

// Campo describes a primitive leaf object in the hierarchy.
type Campo struct {
	gorm.Model `json:"campo"`
	EstudioID  uint
	Value      string `gorm:"column:value" json:"valor"`
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
	gorm.Model `json:"estudio"`
	PacienteID uint         `json:"paciente_id"`
	Campos     []Componente `gorm:"type:Campos" json:"campos"`
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

func initialMigration(db *gorm.DB) {
	db.AutoMigrate(&Estudio{})
	db.AutoMigrate(&Campo{})
}

func main() {

	os.Remove("sqlite3gorm.db")
	db, err := gorm.Open("sqlite3", "sqlite3gorm.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("fallo la conexión. a la concha de tu madre")
	}
	defer db.Close()
	initialMigration(db)

	estudio := NewEstudio(1)

	for i := 0; i < 10; i++ {
		estudio.Add(NewCampo(strconv.Itoa(i)))
		//db.Create(&Campo{Value: strconv.Itoa(i)})
	}
	db.Create(estudio)

	estudio.Traverse()

}
