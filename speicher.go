package main

import "log"

func speicher_1() {
	// speicher_1 START OMIT
	type MyType struct {
		Data string
	}

	x := func() *MyType {
		return &MyType{"some data"}
	}

	b := x()
	log.Println(b.Data)
	// speicher_1 END OMIT
}

func main() {
	speicher_1()
}
