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

// RenderSystem is an ECS system that renders SpriteComponents.
func RenderSystem(entities *[]*Entity, renderer *sdl.Renderer) {
	renderer.Clear()

	for _, entity := range *entities {
		if entity != nil && entity.SpriteComponent != nil {
			renderer.Copy(entity.SpriteComponent.Texture, nil, entity.SpriteComponent.Rect)
		}
	}

	renderer.Present()
}
