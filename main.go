package main

import (
	"awesomeProject/objects"
	"fmt"
	"strings"
)

func main() {

	sStar := "*"
	builder := strings.Builder{}
	iMax := 10

	for i := 0; i < iMax; i++ {
		fmt.Println(builder.String())
		builder.WriteString(sStar)
	}

	job := objects.Job{Bezeichnung: "Anwendungsentwickler", Gehalt: 950, Aufgaben: []string{"Entwickeln", "Support", "Debuggen"}}
	p := objects.Person{Vorname: "Lorenz", Nachname: "E.", Job: job}

	fmt.Println("Aufgaben von " + p.Vorname + " " + p.Nachname)
	for i := 0; i < len(p.Job.Aufgaben); i++ {
		fmt.Println(p.Job.Aufgaben[i])
	}

	testarray := []interface{}{
		32, "Test", map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, map[string]bool{"Mensch": true, "Volljährig": true},
	}

	fmt.Println(dataToJson(testarray))
}
