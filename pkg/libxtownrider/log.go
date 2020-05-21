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
	"fmt"
	"time"
)

const (
	// LogTypeDebug is debug log message type.
	LogTypeDebug = iota
	// LogTypeError is error log message type.
	LogTypeError = iota
)

// Log prints out text with some additions.
func Log(logType int, message string) {
	switch logType {
	case LogTypeDebug:
		fmt.Printf("[%s] [DEBUG] %s\n", time.Now().UTC().Format("2006-01-02T15:04:05Z"), message)
	case LogTypeError:
		fmt.Printf("[%s] [ERROR] %s\n", time.Now().UTC().Format("2006-01-02T15:04:05Z"), message)
	default:
		panic("Incorrect logType")
	}
}
