package main

import (
	"text/template"
	"os"
)

type Persona struct {
	Nombre string
	Edad   int
}

//Devuelve Persona.Nombre solo el .
const tp = `Nombre: {{.Nombre}} Edad : {{.Edad}}`

func main() {

	persona := Persona{"Pablo", 29}

	//Crear nuevo template
	t := template.New("persona")

	//parsear template (mapear)
	t, err := t.Parse(tp)
	if err != nil {
		panic(err)
	}

	err1 := t.Execute(os.Stdout, persona)
	if err1 != nil {
		panic(err)

	}

}
