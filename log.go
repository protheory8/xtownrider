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

package main

import (
	"fmt"
	"time"
)

type logType int

const (
	// logTypeDebug is debug log message type.
	logTypeDebug logType = iota
	// logTypeError is error log message type.
	logTypeError
)

// log prints out text with some additions.
func log(logMessageType logType, message string) {
	switch logMessageType {
	case logTypeDebug:
		fmt.Printf("[%s] [DEBUG] %s\n", time.Now().UTC().Format("2006-01-02T15:04:05Z"), message)
	case logTypeError:
		fmt.Printf("[%s] [ERROR] %s\n", time.Now().UTC().Format("2006-01-02T15:04:05Z"), message)
	default:
		panic("incorrect logMessageType")
	}
}
