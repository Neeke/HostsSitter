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
package Lib

import (
	"flag"
	"github.com/HostsSitter/Enum"
)

type flags struct {
	Command  *string
	SourceId *int
}

var Flags flags

func init() {
	Flags.Command = flag.String("c", Enum.FLAGS_COMMAND_UPDATE, "The Command: update; backup; recovery;")
	Flags.SourceId = flag.Int("sid", Enum.FLAGS_SOURCE_ENUM_360kb, "The SourceId: 1 360kb")

	flag.Parse()
}
