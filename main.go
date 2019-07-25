package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"log"
	"strings"

	client "ddd/dmn-client"
	"ddd/dmn-client/proto"
	"ddd/game"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	initialModeName      = "room1"
	initialModeNamespace = "https://kiegroup.org/dmn/_19909919-4BAA-4792-B7D1-DF492E682883"
	initialText          = "Welcome.\nThis experience is a plain text adventure. Type something..."
)

var (
	user            = game.UserInput{}
	system          = game.SystemOutput{}
	modelName       string
	modelNamespace  string
	backgroundImage *ebiten.Image
	item            string
	enemyHp         int32
	heroHp          int32
	interactions    int32
)

func init() {
	backgroundImage = readImage("background.jpg")
	modelName = initialModeName
	modelNamespace = initialModeNamespace
	enemyHp = 10
	heroHp = 10
}

func main() {
	system.SetText(initialText)

	// ebiten.SetFullscreen(true)
	if err := ebiten.Run(update, 550, 240, 3, "DMN Dungeon Demo"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {

	user.Update()
	system.Update()

	handleMainAction()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	showBackground(screen)
	printUserText(screen)
	printSystemText(screen)
	printHeroHP(screen)
	printEnemyHP(screen)

	return nil
}

func printHeroHP(screen *ebiten.Image) {
	heroHpText := fmt.Sprintf("Hero HP: %d", heroHp)
	ebitenutil.DebugPrintAt(screen, heroHpText, 5, 5)
}

func printEnemyHP(screen *ebiten.Image) {
	// TODO: this logic must exist into the DMN file.
	if interactions >= 5 && modelName == "room1" {
		enemyHpText := fmt.Sprintf("Enemy HP: %d", enemyHp)
		ebitenutil.DebugPrintAt(screen, enemyHpText, 5, 20)
	}
}

func handleMainAction() {

	if isEnterPressed() {

		interactions++

		userInput := getUserInput()

		if isQuit(userInput) {
			os.Exit(0)
		}

		output := client.Evaluate(&proto.GameInput{
			Action:         getAction(userInput),
			Target:         getTarget(userInput),
			Item:           item,
			ModelName:      modelName,
			ModelNamespace: modelNamespace,
			Interactions:   interactions,
			IsHeroTurn:     interactions%2 == 0,
			EnemyHp:        enemyHp,
			HeroHp:         heroHp,
		})

		modelName = output.ModelName
		modelNamespace = output.ModelNamespace
		item = output.NewItem
		enemyHp = output.NewEnemyHp
		heroHp = output.NewHeroHp

		system.SetText(output.Message)

		user.Clear()
	}
}

func getAction(input string) string {
	actions := []string{"attack", "check", "look", "explore", "use"}
	for _, actions := range actions {
		if strings.Contains(input, actions) {
			return actions
		}
	}
	return ""
}

func getTarget(input string) string {
	targets := []string{"room", "chest", "doll", "picture", "papers"}
	for _, target := range targets {
		if strings.Contains(input, target) {
			return target
		}
	}
	return ""
}

func printUserText(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, user.FormattedText, 5, 220)
}

func printSystemText(screen *ebiten.Image) {
	lines := strings.Count(system.Text, "\n")
	ebitenutil.DebugPrintAt(screen, system.FormattedText, 5, 200-(20*lines))
}

func readImage(imagePath string) (image *ebiten.Image) {
	image, _, err := ebitenutil.NewImageFromFile(imagePath, ebiten.FilterDefault)
	if err != nil {
		log.Fatal("Image " + imagePath + " could not be read.")
	}
	return
}

func showBackground(screen *ebiten.Image) {
	imageOptions := ebiten.DrawImageOptions{}
	imageOptions.ColorM.ChangeHSV(30, .1, .1)
	screen.DrawImage(backgroundImage, &imageOptions)
}

func isQuit(action string) bool {
	return action == "quit"
}

func getUserInput() string {
	return strings.ToLower(user.Text[:len(user.Text)-1])
}

func isEnterPressed() bool {
	return strings.Contains(user.Text, "\n") && !system.IsPrinting()
}
