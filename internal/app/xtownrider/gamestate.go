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

import (
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
	gameState.Window = nil
	gameState.Renderer = nil

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
