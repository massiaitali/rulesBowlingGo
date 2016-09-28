package bowling

import (
	"fmt"
	"testing"
)

func scoreChecker(input []Frame, expectedScore int, expectedError error) error {
	score, nil := GetScore(input)
	if score != expectedScore {
		return fmt.Errorf("Score : %d, expected %d", score, expectedScore)
	}
	return nil
}

func NbScore(input []Frame, expectedScore int, expectedError error) error {
	nbTuple := GetNbTuple(input)
	if nbTuple != expectedScore {
		return fmt.Errorf("Nombre de tuples calcules : %d, expected : %d", nbTuple, expectedScore)
	}
	return nil
}

func PositifScore(input []Frame, expectedError error) error {
	positif := NegatifScoreVerif(input)
	if positif == false {
		return fmt.Errorf("Il y a un score negatif !!")
	}
	return nil
}

func MaxScore(input []Frame, expectedError error) error {
	Max := MaxScoreCalc(input)
	if Max == false {
		return fmt.Errorf("Il y a un tuple superieur a 10")
	}
	return nil
}

func StrikeScore(input []Frame, expectedError error) error {
	Strike := StrikeScoreCalc(input)
	if Strike == 0 {
		return fmt.Errorf("Il n'y a pas de strike")
	}
	return nil
}

func SpareScore(input []Frame, expectedError error) error {
	Spare := SpareScoreCalc(input)
	if Spare == 0 {
		return fmt.Errorf("Il n'y a pas de spare")
	}
	return nil
}

func TestScore(t *testing.T) {
	input := []Frame{{6, 4}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 5}, {1, 1}, {5, 3}}
	expected := 98
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestNbTuple(t *testing.T) {
	input := []Frame{{6, 4}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 5}, {1, 1}, {5, 3}}
	expected := 10
	if input[9].firstThrow == 10 {
		expected = 12
	} else if input[9].firstThrow+input[9].secondThrow == 10 {
		expected = 11
	}
	if err := NbScore(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestPositif(t *testing.T) {
	input := []Frame{{6, 4}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 5}, {1, 1}, {5, 3}}
	if err := PositifScore(input, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestMaxScore(t *testing.T) {
	input := []Frame{{6, 4}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 5}, {1, 1}, {5, 3}}
	if err := MaxScore(input, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}
func TestStrike(t *testing.T) {
	input := []Frame{{6, 4}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 5}, {1, 1}, {5, 3}}
	if err := StrikeScore(input, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}
func TestSpare(t *testing.T) {
	input := []Frame{{6, 4}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 0}, {1, 2}, {10, 5}, {1, 1}, {5, 3}}
	if err := SpareScore(input, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}
