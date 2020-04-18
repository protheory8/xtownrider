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

package libxtownrider

import "github.com/veandco/go-sdl2/sdl"

// SpriteComponent is an ECS component that stores texture.
type SpriteComponent struct {
	Texture *sdl.Texture
	Rect    *sdl.Rect
}

func loadTextureFromFile(renderer *sdl.Renderer, file string) (*sdl.Texture, error) {
	var err error
	var surface *sdl.Surface
	var texture *sdl.Texture

	surface, err = sdl.LoadBMP(file)
	if err != nil {
		return nil, err
	}
	defer surface.Free()

	texture, err = renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}

	return texture, nil
}

// NewSpriteComponent makes a new instance of SpriteComponent by reading the BMP file.
func NewSpriteComponent(renderer *sdl.Renderer, file string) (*SpriteComponent, error) {
	var err error
	spriteComponent := new(SpriteComponent)

	spriteComponent.Texture, err = loadTextureFromFile(renderer, file)
	if err != nil {
		return nil, err
	}

	spriteComponent.Rect = new(sdl.Rect)
	_, _, spriteComponent.Rect.W, spriteComponent.Rect.H, err = spriteComponent.Texture.Query()
	if err != nil {
		return nil, err
	}

	return spriteComponent, nil
}
