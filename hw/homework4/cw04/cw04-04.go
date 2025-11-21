/*
  (25 points)

  Follow the instructions below.
  1. Declare a type `named` as a struct with the following fields:
     - name (of type `string`)
  2. Declare a type `person` as a struct composition with `named`
     embedded and with the following additional fields:
     - email (of type `string`)
     - birthYear (of type `int`)
  3. Declare a type `book` as a struct composition with `named`
     embedded and with the following additional fields:
     - author (of type `person`)
     - publishYear (of type `int`)
  4. Declare two different instances of `person`.
  5. Declare two different instances of `book`.
  6. Print the two instances of `book`.
*/

package main

import "fmt"

func main() {
	// --- <code begin> ---
	type named struct {
		name string
	}
	type person struct {
		named
		email     string
		birthYear int
	}
	type book struct {
		named
		author      person
		publishYear int
	}
	s1 := person{named: named{name: "ChungLi"}, email: "teacherWantAnEmail@gmail.com", birthYear: 2007}
	s2 := person{named: named{name: "Luke"}, email: "teacherWantAnotherEmail@gmail.com", birthYear: 2008}
	s3 := book{named: named{name: "Ken"}, author: s1, publishYear: 2010}
	s4 := book{named: named{name: "Ryu"}, author: s2, publishYear: 2011}
	fmt.Println(s3)
	fmt.Println(s4)
	// --- <code end> ---
}
