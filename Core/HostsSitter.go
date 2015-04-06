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
	"fmt"
	"github.com/HostsSitter/Enum"
	"github.com/HostsSitter/Lib"
	"github.com/HostsSitter/Models"
)

func Run() bool {
	if *Lib.Flags.SourceId == Enum.FLAGS_SOURCE_ENUM_360kb {
		Models.Source = Models.Source360
	}

	if *Lib.Flags.Command == Enum.FLAGS_COMMAND_UPDATE {
		return update()
	}

	return true
}

func update() bool {
	Lib.ParseUrl()
	fmt.Println("update ok")
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
