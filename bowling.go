package bowling

type Frame struct {
	firstThrow  int
	secondThrow int
}

func GetScore(game []Frame) (int, error) {
	Strike := StrikeScoreCalc(game)
	Spare := SpareScoreCalc(game)
	score := Strike + Spare
	for i := 0; i < len(game); i++ {
		score = score + game[i].firstThrow + game[i].secondThrow
	}
	return score, nil

}
func GetNbTuple(game []Frame) int {
	score := len(game)
	return score
}
func MaxScoreCalc(game []Frame) bool {
	score := 0
	max := true
	for i := 0; i < len(game); i++ {
		score = game[i].firstThrow + game[i].secondThrow
		if score <= 0 {
			max = false
		}
	}
	return max
}
func NegatifScoreVerif(game []Frame) bool {
	max := true
	for i := 0; i < len(game); i++ {
		if game[i].firstThrow < 0 {
			max = false
		}
		if game[i].secondThrow < 0 {
			max = false
		}
	}
	return max
}

func StrikeScoreCalc(game []Frame) int {
	max := 0
	for i := 0; i < len(game); i++ {
		if game[i].firstThrow == 10 {
			if i+1 == len(game) {
				break
			}
			max = max + game[i+1].firstThrow + game[i+1].secondThrow
		}
	}
	return max
}
func SpareScoreCalc(game []Frame) int {
	max := 0
	for i := 0; i < len(game); i++ {
		if game[i].firstThrow+game[i].secondThrow == 10 {
			if i+1 == len(game) {
				break
			}
			max = max + game[i+1].firstThrow
		}
	}
	return max
}
