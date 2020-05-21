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

import (
	"path/filepath"

	"github.com/veandco/go-sdl2/sdl"
)

// Resource is an interface that all resources must implement.
type Resource interface {
	GetResourceFilename() string
}

// ResourceManager is a struct that manages all your game resources, like
// sprites, audio files, etc.
type ResourceManager struct {
	resources map[string]Resource
}

// NewResourceManager makes new instance of ResourceManager.
func NewResourceManager() ResourceManager {
	resourceManager := ResourceManager{}
	resourceManager.resources = make(map[string]Resource)
	return resourceManager
}

func addResource(resourceManager *ResourceManager, renderer *sdl.Renderer, file string) {
	var err error

	Log(LogTypeDebug, "Adding resource '"+file+"'...")

	switch filepath.Ext(file) {
	case ".bmp":
		var sprite *SpriteComponent
		sprite, err = NewSpriteComponent(renderer, file)
		if err != nil {
			panic(err)
		}

		resourceManager.resources[file] = sprite
	default:
		panic("Unknown resource file extension")
	}
}

// AddResources takes filenames as an input and loads resources into memory.
func (resourceManager *ResourceManager) AddResources(renderer *sdl.Renderer, files []string) {
	for _, file := range files {
		if _, ok := resourceManager.resources[file]; !ok {
			addResource(resourceManager, renderer, file)
		}
	}
}

// Get retrieves resource and returns it.
func (resourceManager *ResourceManager) Get(renderer *sdl.Renderer, file string) Resource {
	if resource, ok := resourceManager.resources[file]; ok {
		return resource
	}

	addResource(resourceManager, renderer, file)
	return resourceManager.resources[file]
}
