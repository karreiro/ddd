package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const prefix = "> "

// UserInput represents the user input into the game
type UserInput struct {
	Text          string
	FormattedText string
	counter       int
}

// Update values during the game loop
func (input *UserInput) Update() {
	input.updateText()
	input.handleBackspaceKey()
	input.handleEnterKey()
	input.updateFormattedText()
}

// Clear deletes the current text
func (input *UserInput) Clear() {
	input.Text = ""
	input.FormattedText = ""
}

func (input *UserInput) updateFormattedText() {
	input.FormattedText = prefix + input.Text + input.getBlinkingCursor()
}

func (input *UserInput) updateText() {
	if len(input.Text) < 48 {
		input.Text += string(ebiten.InputChars())
	}
}

func (input *UserInput) handleEnterKey() {
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
		input.Text += "\n"
	}
}

func (input *UserInput) handleBackspaceKey() {
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(input.Text) >= 1 {
			input.Text = input.Text[:len(input.Text)-1]
		}
	}
}

func (input *UserInput) getBlinkingCursor() string {
	input.counter++
	if input.counter < 30 {
		return "_"
	}
	if input.counter == 60 {
		input.counter = 0
	}
	return ""
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
