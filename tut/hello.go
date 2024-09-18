package main

import "fmt"

func main() {
	fmt.Println("hello world ylo")
	fmt.Printf("Hello, %s!\n", "world2")
	var a int = 2_000_22
	var b string = "sdsds"
	var hindi string = "संकट"
	fmt.Println(a, b)
	fmt.Println((hindi))
	var x int = 10
	var y byte = 100
	var sum3 int = x + int(y)
	var sum4 byte = byte(x) + y
	fmt.Println(sum3, sum4)
	s := []string{"Fred", "Ralph", "Bijou", "Sarah", "Peter"}
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams, s, len(s))
	fmt.Printf("%T\n", s)
}
