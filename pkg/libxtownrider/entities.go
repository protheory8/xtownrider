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

// Entity is simply a collection of ECS components.
type Entity struct {
	SpriteComponent    *SpriteComponent
	TransformComponent *TransformComponent
}

// EntityBuilder is a helper struct that builds entities
type EntityBuilder struct {
	entity Entity
}

// NewEntityBuilder makes a new instance of EntityBuilder.
func NewEntityBuilder() EntityBuilder {
	entityBuilder := EntityBuilder{}
	entityBuilder.entity = Entity{SpriteComponent: nil, TransformComponent: nil}

	return entityBuilder
}

// AddSprite adds sprite to building entity.
func (entityBuilder EntityBuilder) AddSprite(sprite *SpriteComponent) EntityBuilder {
	entityBuilder.entity.SpriteComponent = sprite
	return entityBuilder
}

// SetLocation sets location of the sprite using x and y coordinates.
func (entityBuilder EntityBuilder) SetLocation(x int32, y int32) EntityBuilder {
	if entityBuilder.entity.TransformComponent == nil {
		entityBuilder.entity.TransformComponent = NewTransformComponent(x, y)
		return entityBuilder
	}

	entityBuilder.entity.TransformComponent.X = x
	entityBuilder.entity.TransformComponent.Y = y
	return entityBuilder
}

// Build is a final method to call and it returns built entity.
func (entityBuilder EntityBuilder) Build() Entity {
	return entityBuilder.entity
}

// EntityManager is a helper struct that manages all entities.
type EntityManager struct {
	entities []Entity
}

// NewEntityManager makes a new instance of EntityManager.
func NewEntityManager() EntityManager {
	entityManager := EntityManager{}
	entityManager.entities = []Entity{}

	return entityManager
}

// AddEntity adds entity into the list.
func (entityManager *EntityManager) AddEntity(entity Entity) {
	entityManager.entities = append(entityManager.entities, entity)
}

// GetEntities retrieves entities from EntityManager.
func (entityManager *EntityManager) GetEntities() []Entity {
	return entityManager.entities
}
