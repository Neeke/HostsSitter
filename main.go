/*
  +----------------------------------------------------------------------+
  | HostsSitter                                                          |
  +----------------------------------------------------------------------+
  | This source file is subject to version 2.0 of the Apache license,    |
  | that is bundled with this package in the file LICENSE, and is        |
  | available through the world-wide-web at the following url:           |
  | http://www.apache.org/licenses/LICENSE-2.0.html                      |
  | If you did not receive a copy of the Apache2.0 license and are unable|
  | to obtain it through the world-wide-web, please send a note to       |
  | neeke@php.net so we can mail you a copy immediately.                 |
  +----------------------------------------------------------------------+
  | Author: Neeke.Gao  <neeke@php.net>                                   |
  +----------------------------------------------------------------------+
*/
package main

import (
	"GoCrab/GoCrab/Cli"
	"github.com/HostsSitter/Core"
	"github.com/HostsSitter/Enums"
)

func main() {
	Cli.SetLogger("file", `{"filename":"logs/error.log"}`)

	Cli.CrabApp.Name = Enums.APP_NAME
	Cli.CrabApp.Usage = Enums.APP_USAGE
	Cli.CrabApp.Version = Enums.APP_VERSION

	Cli.CrabApp.Commands = []Cli.Command{
		{
			Name:  "update",
			Usage: "update the hosts use " + Cli.CrabApp.Name,
			Action: func(c *Cli.Context) {
				Core.Run()
			},
		},
	}

	Cli.Info("Agent Running ", Enums.APP_NAME)
	Cli.Run()
}
