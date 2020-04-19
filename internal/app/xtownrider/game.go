// Xtownrider - a car race game written in Go.
// Copyright © 2020 The Xtownrider Contributors
//
// This file is part of Xtownrider.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
