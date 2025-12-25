package main

import (
	"awesomeProject/objects"
	"fmt"
	"strings"
	"unsafe"
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

	response, err := dataToJson(testarray)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(response))
	}

	var unsignedTestInt uint64

	fmt.Println(unsignedTestInt)

	arrayTests()
	fmt.Println("\n\n")

	v := 5
	inc(&v)
	fmt.Println(v)

	fmt.Println("\n\n")

	a := 1
	b := 3
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)

}

func arrayTests() {
	// intArr hat nun 3 Entries Platz mit jeweils 8 Byte. Würde ich [10] nehmen wären es 10*8 = 80 Byte, wieder 8 pro Entry
	// 8 Byte wegen int64, bei int32 wären es 4 Byte, int16 - 2 und int8 - 1
	// würde man int nutzen, dann wäre das Ergebnis Abhängig vom OS, ob es auf 32bit oder 64bit läuft. Für mehr Kontrolle nutzt man int daher nicht.
	intArr := [10]int64{}
	fmt.Printf("Unsafe size: %v", unsafe.Sizeof(intArr))
}

func pointer() {
	x := 10
	y := &x
	// y hält (speichert) nun die adresse x (&x)

	// änderrungen über den Pointer ändert die originalvariable
	*y = 20
	// * hat 2 Bedeutungen. Einmal beim Typ: "Pointer auf..."
	var p *int //nil
	// p ist ein pointer auf einen int
	//oder beim Zugriff "dereferenzieren"
	fmt.Println(*p)
	// lies den Wert, auf den der Pointer zeigt
	fmt.Println(x)
}

func inc(x *int) {
	*x++
}
