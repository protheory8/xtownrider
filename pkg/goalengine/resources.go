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

import (
	"fmt"
	"path/filepath"
	"sync"

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
func NewResourceManager() *ResourceManager {
	resourceManager := new(ResourceManager)
	resourceManager.resources = make(map[string]Resource)
	return resourceManager
}

func (resourceManager *ResourceManager) addResource(id int, wg *sync.WaitGroup, mutex *sync.Mutex, renderer *sdl.Renderer, file string) {
	var err error

	Log(LogTypeDebug, "Adding resource '"+file+"'...")

	switch filepath.Ext(file) {
	case ".bmp":
		var sprite *SpriteComponent
		sprite, err = NewSpriteComponent(renderer, file)
		if err != nil {
			panic(err)
		}

		if mutex != nil {
			fmt.Printf("%d here 1\n", id)
			mutex.Lock()
			fmt.Printf("%d here 2\n", id)
			resourceManager.resources[file] = sprite
			mutex.Unlock()
			fmt.Printf("%d here 3\n", id)
		} else {
			resourceManager.resources[file] = sprite
		}
	default:
		panic("Unknown resource file extension")
	}

	fmt.Printf("%d here 4\n", id)

	if wg != nil {
		fmt.Printf("%d here 5\n", id)
		wg.Done()
		fmt.Printf("%d here 6\n", id)
	}
}

// AddResources takes filenames as an input and loads resources into memory.
func (resourceManager *ResourceManager) AddResources(renderer *sdl.Renderer, files []string) {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for i, file := range files {
		if _, ok := resourceManager.resources[file]; !ok {
			wg.Add(1)
			go resourceManager.addResource(i, &wg, &mutex, renderer, file)
		}
	}

	fmt.Printf("AddResource wait\n")
	wg.Wait()
	fmt.Printf("AddResource wait out\n")
}

// Get retrieves resource and returns it.
func (resourceManager *ResourceManager) Get(renderer *sdl.Renderer, file string) Resource {
	if _, ok := resourceManager.resources[file]; ok {
		return resourceManager.resources[file]
	}

	resourceManager.addResource(0, nil, nil, renderer, file)
	return resourceManager.resources[file]
}
