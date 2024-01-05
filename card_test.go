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

func TestNew(t *testing.T) {
	t.Run("Deck is of the right size", func(t *testing.T) {
		want := 52
		got := len(New())
		if got != want {
			t.Errorf("len(New()) = \"%v\"; want \"%v\"", got, want)
		}
	})
	t.Run("Expect Ace of Diamonds as first Card (DefaultSort)", func(t *testing.T) {
		want := Card{Rank: Ace, Suit: Diamond}
		got := New(DefaultSort)[0]
		if got != want {
			t.Errorf("New(DefaultSort)[0] = \"%v\"; want \"%v\"", got.String(), want.String())
		}
	})
	t.Run("Expect Ace of Diamonds as first Card (SortFunction)", func(t *testing.T) {
		want := Card{Rank: Ace, Suit: Diamond}
		got := New(Sort(Less))[0]
		if got != want {
			t.Errorf("New(Sort(Less))[0] = \"%v\"; want \"%v\"", got.String(), want.String())
		}
	})
	t.Run("Test nbr of Jokers in deck", func(t *testing.T) {
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
	})
	t.Run("No Rank in Deck", func(t *testing.T) {
		deck := New(Filter(func(c Card) bool { return c.Rank == Two || c.Rank == Five }))
		for _, c := range deck {
			if c.Rank == Two || c.Rank == Five {
				t.Errorf("Found %s in Deck, don't want %v", c.String(), []Rank{Two, Five})
			}
		}
	})
	t.Run("No Rank in Deck", func(t *testing.T) {
		want := 4 * 52
		got := len(New(MultipleDecks(4)))
		if got != want {
			t.Errorf("len(New(MultipleDecks(4))) = \"%v\"; want \"%v\"", got, want)
		}
	})
}
