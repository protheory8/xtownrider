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

import "github.com/veandco/go-sdl2/sdl"

type entity struct {
	posX     int32
	posY     int32
	rotation float64
	flip     sdl.RendererFlip

	texture  *sdl.Texture
	textureW int32
	textureH int32
}

func newEntity(texture *sdl.Texture) entity {
	entity := entity{}
	entity.posX = 0
	entity.posY = 0
	entity.rotation = 0.0
	entity.flip = sdl.FLIP_NONE

	if texture != nil {
		var err error
		_, _, entity.textureW, entity.textureH, err = texture.Query()
		if err != nil {
			panic(err)
		}
	}

	entity.texture = texture

	return entity
}
