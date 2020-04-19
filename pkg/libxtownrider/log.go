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
