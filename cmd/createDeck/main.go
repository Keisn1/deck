package main

import (
	"deck"
	"fmt"
)

func main() {
	d := deck.New(deck.WithSorting(deck.DefaultComp))
	for _, c := range d {
		fmt.Println(c.String())
	}

	d = deck.New1(deck.Sort(deck.Less))
	for _, c := range d {
		fmt.Println(c.String())
	}

	d = deck.New1(deck.Shuffle)
	for _, c := range d {
		fmt.Println(c.String())
	}

	d = deck.New1(deck.Shuffle1)
	for _, c := range d {
		fmt.Println(c.String())
	}
}
