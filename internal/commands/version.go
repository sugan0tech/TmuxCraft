// tmuxcraft
// Copyright (c) 2024 sugan0tech
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package commands

import "fmt"

const version = "0.1.0"

func Version() {
  fmt.Println("Tmuxcraft version:", version)
}

