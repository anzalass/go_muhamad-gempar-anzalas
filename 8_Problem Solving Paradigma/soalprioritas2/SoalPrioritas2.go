package soalprioritas2

func Frog(jumps []int) int {
	maxJump := 0
	for i := 0; i < len(jumps); i++ {

		if maxJump+jumps[i] >= len(jumps)-1 {
			maxJump = maxJump + jumps[i]
		}
	}

	return maxJump
}
