//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	// "time"
)

type Suit uint8

const (
	Diamond Suit = iota
	Heart
	Spade
	Club
	Joker
)

var suits = [...]Suit{Diamond, Heart, Spade, Club}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var ranks = [...]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func defaultCards() (cards []Card) {
	for _, rank := range ranks {
		for _, suit := range suits {
			cards = append(cards, Card{
				Suit: suit,
				Rank: rank,
			})
		}
	}
	return
}

func New(optFuncs ...func([]Card) []Card) []Card {
	cards := defaultCards()
	defaultOps := DefaultOpts()

	// for _, opt := range opts {
	// 	cards = opt(cards)
	// }

	cards = Jokers(defaultOps.NbrOfJokers)(cards)
	cards = MultipleDecks(defaultOps.NbrOfDecks)(cards)
	if defaultOps.Shuffle {
		cards = Shuffle(cards)
	}
	for _, rank := range defaultOps.FilterRanks {
		cards = Filter(
			func(c Card) bool { return c.Rank == rank },
		)(cards)
	}

	for _, suit := range defaultOps.FilterSuits {
		cards = Filter(
			func(c Card) bool { return c.Suit == suit },
		)(cards)
	}

	less := func(cards []Card) func(i, j int) bool {
		return func(i, j int) bool {
			return defaultOps.Sort(cards[i]) < defaultOps.Sort(cards[j])
		}
	}

	cards = Sort(less)(cards)

	return cards
}

func applyOpts(cards []Card) []Card {

}

type Opts struct {
	NbrOfJokers int
	NbrOfDecks  int
	Shuffle     bool
	FilterSuits []Suit
	FilterRanks []Rank
	Sort        func(c Card) int
}

func DefaultOpts() Opts {
	return Opts{
		NbrOfJokers: 0,
		NbrOfDecks:  1,
		Shuffle:     false,
		FilterSuits: []Suit{},
		FilterRanks: []Rank{},
		Sort:        absRank,
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func Sort(less func(cards []Card) func(i, j int) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func Shuffle(cards []Card) []Card {
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}

// func Shuffle1(cards []Card) []Card {
// 	ret := make([]Card, len(cards))
// 	r := rand.New(rand.NewSource(time.Now().Unix()))
// 	perm := r.Perm(len(cards))
// 	for i, j := range perm {
// 		ret[i] = cards[j]
// 	}
// 	return ret
// }

func Jokers(nbr int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < nbr; i++ {
			cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
		}
		return cards
	}
}

func checkRankInRanks(givenRank Rank, ranks []Rank) (rankInRanks bool) {
	for _, rank := range ranks {
		if givenRank == rank {
			return true
		}
	}
	return
}

func Filter(filterFunc func(Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var newCards []Card
		for _, c := range cards {
			if !filterFunc(c) {
				newCards = append(newCards, c)
			}
		}
		return newCards
	}
}

func MultipleDecks(nbrOfDecks int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var newDeck []Card
		for i := 0; i < nbrOfDecks; i++ {
			newDeck = append(newDeck, cards...)
		}
		return newDeck
	}
}
