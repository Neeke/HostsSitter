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
package Core

import (
	"GoCrab/GoCrab/Cli"
	"fmt"
	//"github.com/HostsSitter/Enums"
	"github.com/HostsSitter/Lib"
	"github.com/HostsSitter/Models"
)

func Run() bool {
	//if *Lib.Flags.SourceId == Enums.FLAGS_SOURCE_ENUM_360kb {
	Models.Source = Models.Source360
	//}

	return update()
}

func update() bool {
	Lib.ParseUrl()
	Lib.UpdateHosts()
	Cli.Info("update is ok")
	return true
}

func backup() bool {
	fmt.Println("backup ok")
	return true
}

func recovery() bool {
	fmt.Println("recovery ok")
	return true
}
