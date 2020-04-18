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

// Entity is simply a collection of ECS components.
type Entity struct {
	spriteComponent *SpriteComponent
}

// NewEntity makes new instance of Entity
func NewEntity(spriteComponent *SpriteComponent) *Entity {
	entity := new(Entity)

	entity.spriteComponent = spriteComponent

	return entity
}
