package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"t", "s", "w"}
	guess := "s"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s, got %v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"t", "s", "w"}
	guess := "b"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s does not contain letter %s, got %v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")
	if err == nil {
		t.Errorf("Error should be returned when using an invalid word=''")
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "Bob")
	g.MakeAGuess("b")
	validState(t, "goodGuess", g.State)
}

func TestAlreadyGuessed(t *testing.T) {
	g, _ := New(3, "Bob")
	g.MakeAGuess("b")
	g.MakeAGuess("b")
	validState(t, "alreadyGuessed", g.State)
}

func TestGameWon(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	g.MakeAGuess("b")
	validState(t, "won", g.State)
}

func TestGameLost(t *testing.T) {
	g, _ := New(3, "tun")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	g.MakeAGuess("c")
	validState(t, "lost", g.State)
}

func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("state should be '%v', got=%v", expectedState, actualState)
		return false
	}
	return true
}
