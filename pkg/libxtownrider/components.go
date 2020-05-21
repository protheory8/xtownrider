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

package libxtownrider

import "github.com/veandco/go-sdl2/sdl"

// SpriteComponent is an ECS component and resource that stores sprite of the entity.
type SpriteComponent struct {
	filename string
	Texture  *sdl.Texture
	W        int32
	H        int32
}

// NewSpriteComponent makes a new instance of SpriteComponent by reading the BMP file.
func NewSpriteComponent(renderer *sdl.Renderer, file string) (*SpriteComponent, error) {
	var err error
	spriteComponent := new(SpriteComponent)

	spriteComponent.Texture, err = loadTextureFromFile(renderer, file)
	if err != nil {
		return nil, err
	}

	spriteComponent.filename = file

	_, _, spriteComponent.W, spriteComponent.H, err = spriteComponent.Texture.Query()
	if err != nil {
		return nil, err
	}

	return spriteComponent, nil
}

// GetResourceFilename returns the filename of sprite.
func (spriteComponent SpriteComponent) GetResourceFilename() string {
	return spriteComponent.filename
}

// TransformComponent is an ECS component that stores location of the entity.
type TransformComponent struct {
	X int32
	Y int32
}

// NewTransformComponent makes new instance of TransformComponent.
func NewTransformComponent(x int32, y int32) *TransformComponent {
	transformComponent := &TransformComponent{X: x, Y: y}
	return transformComponent
}
