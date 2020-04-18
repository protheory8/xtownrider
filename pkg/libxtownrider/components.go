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
}

// NewSpriteComponent makes a new instance of SpriteComponent by reading the BMP file.
func NewSpriteComponent(renderer *sdl.Renderer, file string) (*SpriteComponent, error) {
	var err error
	var surface *sdl.Surface
	spriteComponent := new(SpriteComponent)

	surface, err = sdl.LoadBMP(file)
	if err != nil {
		return nil, err
	}
	defer surface.Free()

	spriteComponent.Texture, err = renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}

	return spriteComponent, nil
}
