// Copyright 2020 The Xtownrider Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xtownrider

import "github.com/protheory8/xtownrider/pkg/libxtownrider"

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
	entityManager := libxtownrider.NewEntityManager()

	entityManager.AddEntity(libxtownrider.NewEntityBuilder().
		AddSprite(gameState.Renderer, "resources/resource.bmp").
		SetLocation(400, 150).
		Build())

	entityManager.AddEntity(libxtownrider.NewEntityBuilder().
		AddSprite(gameState.Renderer, "resources/resource.bmp").
		SetLocation(100, 150).
		Build())

	entityManager.AddEntity(libxtownrider.NewEntityBuilder().
		AddSprite(gameState.Renderer, "resources/resource.bmp").
		SetLocation(300, 400).
		Build())

	for {
		libxtownrider.RenderSystem(entityManager.GetEntities(), gameState.Renderer)
	}
}
