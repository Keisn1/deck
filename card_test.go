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
		t.Errorf("len(New()) = \"%v\"; want \"%v\"", got, want)
	}
}

func TestDefaultSort(t *testing.T) {
	want := Card{Rank: Ace, Suit: Diamond}
	got := New(WithSort(AbsRank))[0]
	if got != want {
		t.Errorf("New(WithSort(absRank))[0] = \"%v\"; want \"%v\"", got.String(), want.String())
	}
}

func TestReverseSort(t *testing.T) {
	want := Card{Rank: King, Suit: Club}
	got := New(WithReverse())[0]
	if got != want {
		t.Errorf("New(WithReverse())[0] = \"%v\"; want \"%v\"", got.String(), want.String())
	}
}

func TestNbrOfJokers(t *testing.T) {
	want := 10
	deck := New(WithJokers(10))
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

func TestFilterRank(t *testing.T) {
	deck := New(WithFilterRanks([]Rank{Two, Five}))
	for _, c := range deck {
		if c.Rank == Two || c.Rank == Five {
			t.Errorf("Found %s in Deck, don't want %v", c.String(), []Rank{Two, Five})
		}
	}
}

func TestFilterSuit(t *testing.T) {
	deck := New(WithFilterSuits([]Suit{Diamond}))
	for _, c := range deck {
		if c.Suit == Diamond {
			t.Errorf("Found %s in Deck, don't want %v", c.String(), []Suit{Diamond})
		}
	}
}

func TestMultipleDecks(t *testing.T) {
	want := 4 * 52
	got := len(New(WithNbrOfDecks(4)))
	if got != want {
		t.Errorf("len(New1(MultipleDecks(4))) = \"%v\"; want \"%v\"", got, want)
	}
}
