package day7

import (
    _ "embed"
    "fmt"
    "log"
    "regexp"
    "sort"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 7
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day7A",
        Usage: "Day 7 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day7B",
        Usage: "Day 7 part B",
        Action: partB,
    }
)

func partA(ctx *cli.Context) error {
    answer, err := processA(aoc.PreparePuzzle(puzzle))
    if err != nil {
        return err
    }
    log.Printf("Answer A: %v", answer)
    return nil
}

func partB(ctx *cli.Context) error {
    answer, err := processB(aoc.PreparePuzzle(puzzle))
    if err != nil {
        return err
    }
    log.Printf("Answer B: %v", answer)
    return nil
}

type handType int

const (
	FIVE_OF_A_KIND handType = iota
	FOUR_OF_A_KIND
	FULL_HOUSE
	THREE_OF_A_KIND
	TWO_PAIR
	ONE_PAIR
	HIGH_CARD
)

var (
	handTypeString = map[handType]string {
		FIVE_OF_A_KIND: "five of a kind",
		FOUR_OF_A_KIND: "four of a kind",
		FULL_HOUSE: "full house",
		THREE_OF_A_KIND: "three of a kind",
		TWO_PAIR: "two pair",
		ONE_PAIR: "one pair",
		HIGH_CARD: "high card",
	}
)

type hand struct {
	cards string
	bid int
	ht handType
}

func processA(puzzle []string) (int, error) {
	var hands []*hand
	for _, line := range puzzle {
		h, err := decode(line)
		if err != nil {
			return 0, err
		}
		hands = append(hands, h)
		log.Printf("hand: %v", h)
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].ht == hands[j].ht {
			for k:=0; k<len(hands[i].cards); k++ {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}
				return handLess(hands[i].cards[k], hands[j].cards[k])
			}
			return false
		}
		return hands[i].ht < hands[j].ht
	})
	return reportTotal(hands), nil
}

func processB(puzzle []string) (int, error) {
	var hands []*hand
	for _, line := range puzzle {
		h, err := decodeJ(line)
		if err != nil {
			return 0, err
		}
		hands = append(hands, h)
		log.Printf("hand: %v", h)
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].ht == hands[j].ht {
			for k:=0; k<len(hands[i].cards); k++ {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}
				return handLessJ(hands[i].cards[k], hands[j].cards[k])
			}
			return false
		}
		return hands[i].ht < hands[j].ht
	})
	return reportTotal(hands), nil
}

func reportTotal(hands []*hand) int {
	log.Printf("======")
	total := 0
	for i, h := range hands {
		rank := len(hands)-i
		n := h.bid * rank
		log.Printf("#%v: %v (%v * %v = %v)", rank, h, rank, h.bid, n)
		total += n
	}
	return total
}

func decode(line string) (*hand, error) {
	re := regexp.MustCompile(`([A-Z0-9]*) ([0-9]*)`)
	matches := re.FindStringSubmatch(line)
	h := hand{cards: matches[1], bid: aoc.MustAtoi(matches[2])}
	h.calculateType()
	return &h, nil
}

func decodeJ(line string) (*hand, error) {
	re := regexp.MustCompile(`([A-Z0-9]*) ([0-9]*)`)
	matches := re.FindStringSubmatch(line)
	h := hand{cards: matches[1], bid: aoc.MustAtoi(matches[2])}
	m := h.calculateType()
	h.calculateTypeJ(m)
	return &h, nil
}

func (h *hand) String() string {
	hts, ok := handTypeString[h.ht]
	if !ok {
		panic(fmt.Sprintf("%v not in handTypeString map", h.ht))
	}
	return fmt.Sprintf("%v #%v %v", h.cards, h.bid, hts)
}

func (h *hand) calculateType() map[string]int {
	m := make(map[string]int)
	for _, card := range h.cards {
		c := string(card)
		n, ok := m[c]
		if !ok {
			m[c] = 1
		} else {
			m[c] = n+1
		}
	}
	// five of a kind => 5 => len(1)
	// four of a kind => 4, 1 => len(2)
	// full house => 3, 2 => len(2)
	// three of a kind => 3, 1, 1 => len(3)
	// two pair => 2, 2, 1 => len(3)
	// one pair => 2, 1, 1, 1 => len(4)
	// high card => 1, 1, 1, 1, 1 => len(5)
	switch len(m) {
	case 1:
		h.ht = FIVE_OF_A_KIND
	case 2: // four of a kind or full house
		for _, count := range m {
			switch count {
			case 1:
				h.ht = FOUR_OF_A_KIND
			case 2:
				h.ht = FULL_HOUSE
			case 3:
				h.ht = FULL_HOUSE
			case 4:
				h.ht = FOUR_OF_A_KIND
			}
			break
		}
	case 3: // three of a kind or two pair
		for _, count := range m {
			if count == 1 {
				continue
			}
			if count == 2 {
				h.ht = TWO_PAIR
				break
			}
			h.ht = THREE_OF_A_KIND
			break
		}
	case 4:
		h.ht = ONE_PAIR
	case 5:
		h.ht = HIGH_CARD
	default:
	}
	return m
}

var joker = "J"[0]

func (h *hand) calculateTypeJ(m map[string]int) {
	switch h.ht {
	case FIVE_OF_A_KIND:
		// Nothing to do, J can't make a difference.
	case FOUR_OF_A_KIND:
		// If any card is a joker, upgrade to five of a kind.
		// JJJJx -> xxxxx
		// xxxxJ -> xxxxx
		_, ok := m["J"]
		if ok {
			h.ht = FIVE_OF_A_KIND
		}
	case FULL_HOUSE:
		// If any card is a joker, upgrade to five of a kind.
		// JJxxx -> xxxxx
		// xxxJJ -> xxxxx
		_, ok := m["J"]
		if ok {
			h.ht = FIVE_OF_A_KIND
		}
	case THREE_OF_A_KIND:
		// If there's a joker, upgrade to four of a kind.
		// aaabJ -> aaaba
		// aaaJJ -> full house already
		_, ok := m["J"]
		if ok {
			h.ht = FOUR_OF_A_KIND
		}
	case TWO_PAIR:
		// If there's a joker, upgrade to full house or four of a kind.
		// aabbJ -> aabbb
		// aaJJb -> aaaab
		// ktjjt -> ktttt
		n, ok := m["J"]
		if ok {
			if n == 2 {
				h.ht = FOUR_OF_A_KIND
			} else {
				h.ht = FULL_HOUSE
			}
		}
	case ONE_PAIR:
		// If there's a joker, upgrade to three of a kind.
		// Can't be two jokers, otherwise it'd already be two pair.
		// aabcJ -> aabca
		_, ok := m["J"]
		if ok {
			h.ht = THREE_OF_A_KIND
		}
	case HIGH_CARD:
		// If there's a joker, upgrade to one pair.
		// There can't be two jokers otherwise it'd already be one pair.
		// abcdJ -> abcdd
		_, ok := m["J"]
		if ok {
			h.ht = ONE_PAIR
		}
	}
}

var (
	cardRank = map[byte]int{
		'A': 1,
		'K': 2,
		'Q': 3,
		'J': 4,
		'T': 5,
		'9': 6,
		'8': 7,
		'7': 8,
		'6': 9,
		'5': 10,
		'4': 11,
		'3': 12,
		'2': 13,
	}
)

func handLess(a, b byte) bool {
	ca, ok := cardRank[a]
	if !ok {
		panic(fmt.Sprintf("%v not in card rank", b))
	}
	cb, ok := cardRank[b]
	if !ok {
		panic(fmt.Sprintf("%v not in card rank", b))
	}
	return ca < cb
}

func handLessJ(a, b byte) bool {
	ca, ok := cardRank[a]
	if !ok {
		panic(fmt.Sprintf("%v not in card rank", b))
	}
	cb, ok := cardRank[b]
	if !ok {
		panic(fmt.Sprintf("%v not in card rank", b))
	}
	if ca == 4 { return false }
	if cb == 4 { return true }
	return ca < cb
}

func testHL() {
	s := "AKQJT98765432"
	t := "3"[0]
	for i:=0; i<len(s); i++ {
		log.Printf("%v<%v = %v", string(t), string(s[i]), handLess(t, s[i]))
	}
}
