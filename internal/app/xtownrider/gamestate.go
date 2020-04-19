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

import (
	"github.com/protheory8/xtownrider/pkg/libxtownrider"
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
	libxtownrider.Log(libxtownrider.LogTypeDebug, "Dropping GameState...")

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
