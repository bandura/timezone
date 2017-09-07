package timezone

import (
	"testing"
	"fmt"
	"log"
)

func TestExists(t *testing.T) {
	ok := Exists("America/Chicago")
	if !ok {
		t.Fatal("Test exists failed")
	}
	ok = Exists("America/Columbia")
	if ok {
		t.Fatal("Test not exists failed")
	}
}

func TestOffsetText(t *testing.T) {
	s := OffsetText("America/Chicago")
	fmt.Printf(s)
	if s != "-06:00" {
		t.Fatal("Failed OffsetText")
	}
}

func TestOffset(t *testing.T) {
	n, ok := Offset("America/Chicago")
	fmt.Println(n)
	if !ok {
		t.Fatal("Failed good offset test")
	}
	n, ok = Offset("Australia/Eucla")
	fmt.Println(n)
	if !ok {
		t.Fatal("Failed good offset test")
	}
	n, ok = Offset("America/Paris")
	fmt.Println(n)
	if ok {
		t.Fatal("Failed bad offset test")
	}
}

func TestNames(t *testing.T) {
	a := Names()
	log.Printf("%#q\n", a)
}

func TestEntries(t *testing.T) {
	e := Entries()
	log.Printf("%v\n", e)
}
