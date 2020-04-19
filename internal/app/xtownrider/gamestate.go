package xtownrider

import (
	"github.com/protheory8/goalengine"
	"github.com/veandco/go-sdl2/sdl"
)

// GameState is a struct that stores current state of the game.
type GameState struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
}

// NewGameState makes a new instance of GameState struct.
func NewGameState() *GameState {
	var err error
	gameState := new(GameState)

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		gameState.Drop()
		panic(err)
	}

	gameState.Window, err = sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WindowWidth, WindowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		gameState.Drop()
		panic(err)
	}

	gameState.Renderer, err = sdl.CreateRenderer(gameState.Window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		gameState.Drop()
		panic(err)
	}

	return gameState
}

// Drop invokes Destroy on every SDL component.
func (gameState *GameState) Drop() {
	goalengine.Log(goalengine.LogTypeDebug, "Dropping GameState...")

	if gameState.Window != nil {
		gameState.Window.Destroy()
		gameState.Window = nil
	}

	if gameState.Renderer != nil {
		gameState.Renderer.Destroy()
		gameState.Renderer = nil
	}

	sdl.Quit()
}
