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
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// gameState is a struct that stores current state of the game.
type gameState struct {
	shouldQuit bool
	window     *sdl.Window
	renderer   *sdl.Renderer
}

// newGameState makes a new instance of gameState struct.
func newGameState() gameState {
	var err error
	gameState := gameState{}

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		gameState.drop()
		panic(err)
	}

	err = img.Init(img.INIT_PNG)
	if err != nil {
		gameState.drop()
		panic(err)
	}

	gameState.window, err = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		gameState.drop()
		panic(err)
	}

	gameState.renderer, err = sdl.CreateRenderer(gameState.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		gameState.drop()
		panic(err)
	}

	return gameState
}

// drop invokes Destroy on every SDL component.
func (gameState *gameState) drop() {
	log(logTypeDebug, "Dropping gameState...")

	if gameState.window != nil {
		gameState.window.Destroy()
		gameState.window = nil
	}

	if gameState.renderer != nil {
		gameState.renderer.Destroy()
		gameState.renderer = nil
	}

	sdl.Quit()
}
