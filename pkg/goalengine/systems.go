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

package goalengine

import "github.com/veandco/go-sdl2/sdl"

// RenderSystem is an ECS system that renders SpriteComponents.
func RenderSystem(entities *[]*Entity, renderer *sdl.Renderer) {
	renderer.Clear()

	for _, entity := range *entities {
		if entity != nil && entity.SpriteComponent != nil {
			renderer.Copy(entity.SpriteComponent.Texture, nil, &sdl.Rect{X: entity.TransformComponent.X, Y: entity.TransformComponent.Y,
				W: entity.SpriteComponent.W, H: entity.SpriteComponent.H})
		}
	}

	renderer.Present()
}

// FreeSystem deallocated and destroys everything.
func FreeSystem(entities *[]*Entity) {
	for _, entity := range *entities {
		if entity != nil {
			if entity.SpriteComponent != nil && entity.SpriteComponent.Texture != nil {
				entity.SpriteComponent.Texture.Destroy()
			}
		}
	}
}