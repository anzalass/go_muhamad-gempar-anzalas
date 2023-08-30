package soalprioritas2

func PlayingDomino(cards [][]int, deck []int) interface{} {
	for _, card := range cards {
		if card[0] == deck[1] || card[1] == deck[0] {
			return card
		}
	}

	return "Tidak Ada"
}
