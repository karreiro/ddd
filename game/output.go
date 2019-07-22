package game

var counter = 0

// SystemOutput represents the game output into the game
type SystemOutput struct {
	Text          string
	FormattedText string
	textPosition  int
	counter       int
}

// SetText sets a new text into the output
func (output *SystemOutput) SetText(text string) {
	output.Text = text
	output.textPosition = 0
}

// Update values during the game loop
func (output *SystemOutput) Update() {
	if output.IsPrinting() {
		output.updateFormattedText()
	}
}

// Clear deletes the current text
func (output *SystemOutput) Clear() {
	output.Text = ""
	output.FormattedText = ""
	output.textPosition = 0
}

// IsPrinting checks if the system is printing a output
func (output *SystemOutput) IsPrinting() bool {
	return output.textPosition < len(output.Text)
}

func (output *SystemOutput) updateFormattedText() {

	interval := 1
	output.counter++

	if output.counter == interval {
		output.textPosition++
		output.FormattedText = output.Text[:output.textPosition]
	}
	if output.counter > interval {
		output.counter = 0
	}
}
