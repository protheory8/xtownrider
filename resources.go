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

package xtownrider

import (
	"path/filepath"

	"github.com/veandco/go-sdl2/sdl"
)

// resourceManager is a struct that manages all your game resources, like
// sprites, audio files, etc.
type resourceManager struct {
	resources map[string]interface{}
}

// newResourceManager makes new instance of resourceManager.
func newResourceManager() resourceManager {
	resourceManager := resourceManager{}
	resourceManager.resources = make(map[string]interface{})
	return resourceManager
}

func addResource(resourceManager *resourceManager, renderer *sdl.Renderer, fileName string) {
	var err error

	log(logTypeDebug, "Adding resource '"+fileName+"'...")

	switch filepath.Ext(fileName) {
	case ".bmp":
		var sprite *sdl.Texture
		sprite, err = loadBmpTextureFromFile(renderer, fileName)
		if err != nil {
			panic(err)
		}

		resourceManager.resources[fileName] = sprite
	default:
		panic("unknown resource file extension")
	}
}

// addResources takes filenames as an input and loads resources into memory.
func (resourceManager *resourceManager) addResources(renderer *sdl.Renderer, fileNames []string) {
	for _, file := range fileNames {
		if _, ok := resourceManager.resources[file]; !ok {
			addResource(resourceManager, renderer, file)
		}
	}
}

// get retrieves resource and returns it.
func (resourceManager *resourceManager) get(renderer *sdl.Renderer, file string) interface{} {
	if _, ok := resourceManager.resources[file]; !ok {
		addResource(resourceManager, renderer, file)
	}

	return resourceManager.resources[file]
}
