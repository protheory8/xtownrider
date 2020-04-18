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

// Entity is simply a collection of ECS components.
type Entity struct {
	SpriteComponent *SpriteComponent
}

// NewEntity makes new instance of Entity
func NewEntity(x int32, y int32, spriteComponent *SpriteComponent) *Entity {
	entity := new(Entity)

	entity.SpriteComponent = spriteComponent

	return entity
}

// EntityBuilder is a helper struct that builds entities
type EntityBuilder struct {
	entity *Entity
}

// NewEntityBuilder makes a new instance of EntityBuilder.
func NewEntityBuilder() *EntityBuilder {
	entityBuilder := new(EntityBuilder)

	entityBuilder.entity = new(Entity)
	entityBuilder.entity.SpriteComponent = nil

	return entityBuilder
}

// AddSprite adds sprite to building entity.
func (entityBuilder EntityBuilder) AddSprite(renderer *sdl.Renderer, file string) EntityBuilder {
	var err error

	entityBuilder.entity.SpriteComponent, err = NewSpriteComponent(renderer, file)
	if err != nil {
		panic(err)
	}

	_, _, entityBuilder.entity.SpriteComponent.Rect.W, entityBuilder.entity.SpriteComponent.Rect.H, err = entityBuilder.entity.SpriteComponent.Texture.Query()
	if err != nil {
		panic(err)
	}

	return entityBuilder
}

// SetLocation sets location of the sprite using x and y coordinates.
func (entityBuilder EntityBuilder) SetLocation(x int32, y int32) EntityBuilder {
	if entityBuilder.entity.SpriteComponent == nil {
		panic("Sprite wasn't added yet")
	}

	entityBuilder.entity.SpriteComponent.Rect.X = x
	entityBuilder.entity.SpriteComponent.Rect.Y = y

	return entityBuilder
}

// Build is a final method to call and it returns built entity.
func (entityBuilder EntityBuilder) Build() *Entity {
	return entityBuilder.entity
}

// EntityManager is a helper struct that manages all entities.
type EntityManager struct {
	entities []*Entity
}

// NewEntityManager makes a new instance of EntityManager.
func NewEntityManager() *EntityManager {
	entityManager := new(EntityManager)

	entityManager.entities = []*Entity{}

	return entityManager
}

// AddEntity adds entity into the list.
func (entityManager *EntityManager) AddEntity(entity *Entity) {
	entityManager.entities = append(entityManager.entities, entity)
}

// GetEntities retrieves entities from EntityManager.
func (entityManager *EntityManager) GetEntities() *[]*Entity {
	return &entityManager.entities
}
