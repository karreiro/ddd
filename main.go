package main

import (
	_ "image/jpeg"

	"log"
	"strings"

	"ddd/game"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	user    = game.UserInput{}
	system  = game.SystemOutput{}
	counter = 0
	img     *ebiten.Image
)

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("scenario.jpg", ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	system.SetText("What's your name?\nBe, careful you won't be able to change it.")

	ebiten.SetFullscreen(true)
	if err := ebiten.Run(update, 430, 240, 3, "DMN Dungeon Demo"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {

	user.Update()
	system.Update()

	if strings.Contains(user.Text, "\n") && !system.IsPrinting() {
		system.SetText("Hi, " + user.Text[:len(user.Text)-1])
		user.Clear()
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	lines := strings.Count(system.FormattedText, "\n")
	imageOptions := ebiten.DrawImageOptions{}
	imageOptions.ColorM.ChangeHSV(30, .1, .1)

	screen.DrawImage(img, &imageOptions)
	ebitenutil.DebugPrintAt(screen, system.FormattedText, 5, 200-(20*lines))
	ebitenutil.DebugPrintAt(screen, user.FormattedText, 5, 220)
	return nil
}
