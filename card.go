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

type Opts struct {
	rankFuncForSort func(c Card) int
	Reverse         bool
	NbrOfJokers     int
	NbrOfDecks      int
	SuitsToFilter   []Suit
	RanksToFilter   []Rank
	WithShuffle     bool
}

func getDefaultOpts() Opts {
	return Opts{
		rankFuncForSort: AbsRank,
		Reverse:         false,
		NbrOfJokers:     0,
		NbrOfDecks:      1,
		WithShuffle:     false,
		SuitsToFilter:   []Suit{},
		RanksToFilter:   []Rank{},
	}
}

type OptFunc func(*Opts)

func New(optFuncs ...OptFunc) []Card {
	cards := getDefaultCards()
	opts := getDefaultOpts()

	for _, optFunc := range optFuncs {
		optFunc(&opts)
	}

	cards = applyOpts(cards, opts)
	return cards
}

func WithSort(rankFuncForSort func(c Card) int) OptFunc {
	return func(o *Opts) {
		o.rankFuncForSort = rankFuncForSort
	}
}

func WithReverse() OptFunc {
	return func(o *Opts) {
		o.Reverse = true
	}
}

func WithJokers(nbrOfJokers int) OptFunc {
	return func(o *Opts) {
		o.NbrOfJokers = nbrOfJokers
	}
}

func WithNbrOfDecks(nbrOfDecks int) OptFunc {
	return func(o *Opts) {
		o.NbrOfDecks = nbrOfDecks
	}
}

func WithFilterSuits(suitsToFilter []Suit) OptFunc {
	return func(o *Opts) {
		o.SuitsToFilter = suitsToFilter
	}
}

func WithFilterRanks(ranksToFilter []Rank) OptFunc {
	return func(o *Opts) {
		o.RanksToFilter = ranksToFilter
	}
}

func WithShuffle(withShuffle bool) OptFunc {
	return func(o *Opts) {
		o.WithShuffle = withShuffle
	}
}

func WithMultipleDecks(nbrOfDecks int) OptFunc {
	return func(o *Opts) {
		o.NbrOfDecks = nbrOfDecks
	}
}

func getDefaultCards() (cards []Card) {
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

func applyOpts(cards []Card, opts Opts) []Card {
	cards = applyJokers(opts.NbrOfJokers)(cards)
	cards = applyMultipleDecks(opts.NbrOfDecks)(cards)

	if opts.WithShuffle {
		cards = applyShuffle(cards)
	}

	for _, rank := range opts.RanksToFilter {
		cards = applyFilter(
			func(c Card) bool { return c.Rank == rank },
		)(cards)
	}

	for _, suit := range opts.SuitsToFilter {
		cards = applyFilter(
			func(c Card) bool { return c.Suit == suit },
		)(cards)
	}

	comp := getCompFunc(opts)
	sort.Slice(cards, comp(cards))
	return cards
}

func getCompFunc(opts Opts) func(cards []Card) func(i, j int) bool {
	return func(cards []Card) func(i, j int) bool {
		return func(i, j int) bool {
			if opts.Reverse {
				return opts.rankFuncForSort(cards[i]) > opts.rankFuncForSort(cards[j])
			}
			return opts.rankFuncForSort(cards[i]) < opts.rankFuncForSort(cards[j])
		}
	}
}

func AbsRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func applyJokers(nbr int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < nbr; i++ {
			cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
		}
		return cards
	}
}

func applyMultipleDecks(nbrOfDecks int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var newDeck []Card
		for i := 0; i < nbrOfDecks; i++ {
			newDeck = append(newDeck, cards...)
		}
		return newDeck
	}
}

func applyFilter(filterFunc func(Card) bool) func([]Card) []Card {
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

func applyShuffle(cards []Card) []Card {
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}
