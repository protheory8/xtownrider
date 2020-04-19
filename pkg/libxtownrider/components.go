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
