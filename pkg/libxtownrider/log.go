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

import (
	"fmt"
	"time"
)

const (
	// LogTypeDebug is debug log message type.
	LogTypeDebug = iota
	// LogTypeError is debug log message type.
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
