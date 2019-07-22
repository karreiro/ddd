package main

import (
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
	initialText          = `Welcome.
This experience is plain text adventure.
You are a warrior.
You can LOOK at things.
You can ATACK anything with your sword.
You can OPEN a chest.
During this adventure you will learn new skills. Be careful.`
)

var (
	user             = game.UserInput{}
	system           = game.SystemOutput{}
	pendingAction    = func() {}
	hasPendingAction bool
	modelName        string
	modelNamespace   string
	monster          bool
	backgroundImage  *ebiten.Image
	userItem         string
	userName         string
	tableItem        string
)

func init() {
	backgroundImage = readImage("background.jpg")
	modelName = initialModeName
	modelNamespace = initialModeNamespace
}

func main() {
	system.SetText(initialText)

	// ebiten.SetFullscreen(true)
	if err := ebiten.Run(update, 430, 240, 3, "DMN Dungeon Demo"); err != nil {
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

	return nil
}

func handleMainAction() {

	if isEnterPressed() {

		userInput := getUserInput()

		if isQuit(userInput) {
			os.Exit(0)
		}

		output := client.Evaluate(&proto.GameInput{
			Action:         getAction(userInput),
			Target:         getTarget(userInput),
			Item:           userItem,
			ModelName:      modelName,
			ModelNamespace: modelNamespace,
		})
		modelName = output.ModelName
		modelNamespace = output.ModelNamespace

		if hasNewItem(output) {
			handleNewItem(output)
		} else {
			system.SetText(output.Message)
		}

		user.Clear()
	}
}

func handlePendingAction() {
	userInput := getUserInput()
	if isEnterPressed() {
		if strings.Contains(userInput, "ye") {
			pendingAction()
		} else {
			system.SetText("OK.")
		}
		user.Clear()
	}
}

func getAction(input string) string {
	if strings.Contains(input, "look") {
		return "look"
	}
	if strings.Contains(input, "attack") {
		return "attack"
	}
	if strings.Contains(input, "run") {
		return "run"
	}
	if strings.Contains(input, "use") {
		return "use"
	}
	return input
}

func getTarget(input string) string {
	if strings.Contains(input, "chest") {
		return "chest"
	}
	if strings.Contains(input, "enemy") {
		return "enemy"
	}
	if strings.Contains(input, "floor") {
		return "floor"
	}
	if strings.Contains(input, "key") {
		return userItem
	}
	return input
}

func hasNewItem(output *proto.GameOutput) bool {
	return len(output.GetItem()) > 0
}

func handleNewItem(output *proto.GameOutput) {

	message := output.Message
	tableItem = output.GetItem()

	if len(tableItem) > 0 {
		if len(userItem) > 0 {
			message += "\nYou are holding a " + userItem + "."
			message += "\nDo you wan to leave it and take the " + tableItem + "? (Yes/No)"
			hasPendingAction = true
			pendingAction = func() {
				userItem = tableItem
				hasPendingAction = false
				system.SetText("Ok. Now you have the " + userItem + ". What are you going to do?")
			}
		} else {
			userItem = tableItem
		}
	}

	system.SetText(message)
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
