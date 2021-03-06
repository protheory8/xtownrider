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

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowTitle  = "Xtownrider"
	windowWidth  = 800
	windowHeight = 600
)

// gameInit runs some initialization functions.
func gameInit() (gameState, resourceManager) {
	gameState := newGameState()
	resourceManager := newResourceManager()

	resourceManager.addResources(gameState.renderer, []string{"resources/car.png"})
	gameState.entities = append(gameState.entities, newEntity(resourceManager.get(gameState.renderer, "resources/car.png").(*sdl.Texture)))
	gameState.entities[0].posX = 30
	gameState.entities[0].posY = 300
	gameState.entities[0].flip = sdl.FLIP_HORIZONTAL

	return gameState, resourceManager
}

// mainGameLoop is the main game loop of the game.
// It mainly does three things:
// 1. Updates the game state.
// 2. Renders to screen.
// 3. Handles input.
func mainGameLoop(gameState *gameState, resourceManager *resourceManager) {
	for !gameState.shouldQuit {
		update(gameState)
		render(gameState)
		handleInput(gameState)
	}
}

func update(gameState *gameState) {
	gameState.entities[0].posX++
}

func render(gameState *gameState) {
	gameState.renderer.Clear()

	for _, entity := range gameState.entities {
		if entity.texture != nil {
			gameState.renderer.CopyEx(entity.texture, nil, &sdl.Rect{X: entity.posX, Y: entity.posY,
				W: entity.textureW, H: entity.textureH}, entity.rotation, nil, entity.flip)
		}
	}

	gameState.renderer.Present()
}

func handleInput(gameState *gameState) {
	var event sdl.Event

	for {
		event = sdl.PollEvent()
		if event == nil {
			return
		}

		switch event.(type) {
		case *sdl.QuitEvent:
			gameState.shouldQuit = true
		}
	}
}
