package main

import (
	"fmt"
	"sort"
)

//- Partindo do código abaixo, ordene os []user por idade e sobrenome.
//- https://play.golang.org/p/BVRZTdlUZ_

//- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type byLastName []user

func (l byLastName) Len() int {
	return len(l)
}

func (l byLastName) Less(i, j int) bool {
	return l[i].Last < l[j].Last
}

func (l byLastName) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users1 := []user{u1, u2, u3}

	fmt.Println(users1, "\n")

	sort.Sort(byAge(users1))
	fmt.Println(users1, "\n")

	sort.Sort(byLastName(users1))
	fmt.Println(users1, "\n")

	for _, v := range users1 {
		sort.Strings(v.Sayings)
	}

	for _, v := range users1 {
		fmt.Println("\n", v.First)
		for i, value := range v.Sayings {
			fmt.Println("\t", i, ".", value)
		}
	}

	// your code goes here
}
