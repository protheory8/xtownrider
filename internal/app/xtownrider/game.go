package xtownrider

import "github.com/protheory8/goalengine"

const (
	// WindowTitle is a constant that stores the name of the window.
	WindowTitle = "Xtownrider"
	// WindowWidth is a constant that stores the width of the window.
	WindowWidth = 800
	// WindowHeight is a constant that stores the height of the window.
	WindowHeight = 600
)

// MainGameLoop is the main game loop of the game
func MainGameLoop() {
	gameState := NewGameState()
	defer gameState.Drop()
	entityManager := goalengine.NewEntityManager()

	entityManager.AddEntity(goalengine.NewEntityBuilder().
		AddSprite(gameState.Renderer, "resources/resource.bmp").
		SetLocation(400, 150).
		Build())

	entityManager.AddEntity(goalengine.NewEntityBuilder().
		AddSprite(gameState.Renderer, "resources/resource.bmp").
		SetLocation(100, 150).
		Build())

	entityManager.AddEntity(goalengine.NewEntityBuilder().
		AddSprite(gameState.Renderer, "resources/resource.bmp").
		SetLocation(300, 400).
		Build())

	for {
		goalengine.RenderSystem(entityManager.GetEntities(), gameState.Renderer)
	}
}
