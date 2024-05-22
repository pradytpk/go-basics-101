package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	lkt bool
}

func main() {
	p1 := person{
		first: "Test",
		last:  "kuamr",
		age:   32,
	}
	fmt.Println(p1)
	fmt.Printf("%T \t %#v", p1, p1)

	sa1 := secretAgent{
		person: person{first: "Test1",
			last: "kuamr1",
			age:  321},
		lkt: true,
	}

	fmt.Println(sa1)
	fmt.Printf("%T \t %#v", sa1, sa1)

	// anonymous struct
	p2 := struct {
		first string
		last  string
	}{
		first: "testan",
		last:  "one",
	}

	fmt.Println(p2)
	fmt.Printf("%T \t %v", p2, p2)
}
