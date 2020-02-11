package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Paciente struct {
	Nombre   string `json:"nombre"`
	Estudios []Estudio
}

type Estudio struct {
	CampoA string `json:campo_a`
	CampoB string `json:campo_b`
	CampoC string `json:campo_c`
}

type Visita struct {
	Fecha    time.Time `json:fecha`
	Paciente Paciente
}

func main() {
	p := Paciente{
		Nombre: "Juan Jose Paso",
		Estudios: []Estudio{
			Estudio{
				CampoA: "Bien",
				CampoB: "Bien",
				CampoC: "Bien",
			},
			Estudio{
				CampoA: "Para el orto",
				CampoB: "Se murio",
				CampoC: "cago fuego",
			},
		},
	}
	v := Visita{
		Paciente: p,
		Fecha:    time.Now(),
	}
	js, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))
}
