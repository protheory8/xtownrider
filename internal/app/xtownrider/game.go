// This file is part of Xtownrider.
//
// Copyright © 2020 The Xtownrider Contributors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package xtownrider

import (
	"github.com/protheory8/xtownrider/pkg/libxtownrider"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// WindowTitle is a constant that stores the name of the window.
	WindowTitle = "Xtownrider"
	// WindowWidth is a constant that stores the width of the window.
	WindowWidth = 800
	// WindowHeight is a constant that stores the height of the window.
	WindowHeight = 600
)

// Init runs some initialization functions.
func Init() (GameState, libxtownrider.ResourceManager, libxtownrider.EntityManager) {
	gameState := NewGameState()
	resourceManager := libxtownrider.NewResourceManager()
	entityManager := libxtownrider.NewEntityManager()

	resourceManager.AddResources(gameState.Renderer, []string{"resources/resource.bmp", "resources/resource2.bmp"})

	entityManager.AddEntity(libxtownrider.NewEntityBuilder().
		AddSprite(resourceManager.Get(gameState.Renderer, "resources/resource.bmp").(*libxtownrider.SpriteComponent)).
		SetLocation(400, 150).
		Build())

	entityManager.AddEntity(libxtownrider.NewEntityBuilder().
		AddSprite(resourceManager.Get(gameState.Renderer, "resources/resource2.bmp").(*libxtownrider.SpriteComponent)).
		SetLocation(100, 150).
		Build())

	entityManager.AddEntity(libxtownrider.NewEntityBuilder().
		AddSprite(resourceManager.Get(gameState.Renderer, "resources/resource.bmp").(*libxtownrider.SpriteComponent)).
		SetLocation(300, 400).
		Build())

	return gameState, resourceManager, entityManager
}

// MainGameLoop is the main game loop of the game.
// It mainly does three things:
// 1. Updates the game state.
// 2. Renders to screen.
// 3. Handles input.
func MainGameLoop(gameState *GameState, resourceManager *libxtownrider.ResourceManager, entityManager *libxtownrider.EntityManager) {
	for !gameState.ShouldQuit {
		update(gameState)
		render(gameState, entityManager.GetEntities())
		handleInput(gameState)
	}
}

func update(_ *GameState) {}

func render(gameState *GameState, entities []libxtownrider.Entity) {
	libxtownrider.RenderSystem(entities, gameState.Renderer)
}

func handleInput(gameState *GameState) {
	var event sdl.Event
	var eventType uint32

	for {
		event = sdl.PollEvent()
		if event == nil {
			return
		}

		eventType = event.GetType()
		switch eventType {
		case sdl.QUIT:
			gameState.ShouldQuit = true
		}
	}
}
