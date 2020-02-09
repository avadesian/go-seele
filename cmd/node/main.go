/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package main

import (
	"github.com/seeleteam/go-seele/cmd/node/cmd"
	"github.com/stackimpact/stackimpact-go"
)

func main() {
	stackimpact.Start(stackimpact.Options{
			AgentKey: "2c8a8111619d866c911a28ef8a6bd1eaf872dc58",
		AppName: "MyGoApp",
		})
	cmd.Execute()
}
