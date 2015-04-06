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
package Models

import (
	"github.com/HostsSitter/Enum"
	"runtime"
)

type hosts struct {
	OS          string
	HostPath    string
	HostPreStr  string
	HostPostStr string
}

var Hosts hosts

func init() {
	Hosts.OS = runtime.GOOS
	if Hosts.OS == "windows" {
		Hosts.HostPath = Enum.HOST_PATH_LINUX
	} else {
		Hosts.HostPath = Enum.HOST_PATH_WINDOWS
	}

	Hosts.HostPreStr = Enum.HOST_PRE_STR
	Hosts.HostPostStr = Enum.HOST_POST_STR
}
