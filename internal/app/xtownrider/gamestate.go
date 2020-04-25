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
	"github.com/protheory8/goalengine"
	"github.com/veandco/go-sdl2/sdl"
)

// GameState is a struct that stores current state of the game.
type GameState struct {
	ShouldQuit bool
	Window     *sdl.Window
	Renderer   *sdl.Renderer
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
