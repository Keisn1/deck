package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	c := Card{Rank: Ace, Suit: Diamond}
	fmt.Println(c.String())

	// Output:
	// Ace of Diamonds
}

func TestLen(t *testing.T) {
	want := 52
	got := len(New())
	if got != want {
		t.Errorf("len(New1()) = \"%v\"; want \"%v\"", got, want)
	}
}

func TestDefaultSort(t *testing.T) {
	want := Card{Rank: Ace, Suit: Diamond}
	got := New(DefaultSort)[0]
	if got != want {
		t.Errorf("New1(DefaultSort)[0] = \"%v\"; want \"%v\"", got.String(), want.String())
	}
}

func TestSortLess(t *testing.T) {
	want := Card{Rank: Ace, Suit: Diamond}
	got := New(Sort(Less))[0]
	if got != want {
		t.Errorf("New1(Sort(Less))[0] = \"%v\"; want \"%v\"", got.String(), want.String())
	}
}

func TestNbrOfJokers(t *testing.T) {
	want := 10
	deck := New(Jokers(10))
	count := 0
	for _, c := range deck {
		if c.Suit == Joker {
			count++
		}
	}
	if count != want {
		t.Errorf("Number of Jokers in Deck = \"%v\"; want \"%v\"", count, want)
	}
}

func TestFilter(t *testing.T) {
	deck := New(Filter(func(c Card) bool { return c.Rank == Two || c.Rank == Five }))
	for _, c := range deck {
		if c.Rank == Two || c.Rank == Five {
			t.Errorf("Found %s in Deck, don't want %v", c.String(), []Rank{Two, Five})
		}
	}
}

func TestMultipleDecks(t *testing.T) {
	want := 4 * 52
	got := len(New(MultipleDecks(4)))
	if got != want {
		t.Errorf("len(New1(MultipleDecks(4))) = \"%v\"; want \"%v\"", got, want)
	}
}
